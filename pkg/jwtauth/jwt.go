package jwtauth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/logger"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"openiam/app/system/models"
	"time"
)

type Claims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	jti string
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

func GenerateTokens(userId string, username string) (result *TokenPair, err error) {
	var (
		now                         = time.Now()
		jti                         = uuid.New().String()
		issuer                      = viper.GetString("jwt.issuer")
		accessToken, refreshToken   *jwt.Token
		refreshClaims               *RefreshClaims
		signedAccess, signedRefresh string
	)

	// Access Token
	accessClaims := &Claims{
		UserId:   userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(viper.GetInt("jwt.accessToken.expires")) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    issuer,
			ID:        jti,
		},
	}

	accessToken = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccess, err = accessToken.SignedString(viper.GetString("jwt.accessToken.secret"))
	if err != nil {
		logger.Errorf("jwt sign access token err: %v", err)
		return
	}

	// Refresh Token（不包含用户敏感信息）
	refreshClaims = &RefreshClaims{
		jti: jti,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(viper.GetInt("jwt.refreshToken.expires")) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    issuer,
			ID:        uuid.New().String(),
		},
	}

	refreshToken = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefresh, err = refreshToken.SignedString(viper.GetString("jwt.accessToken.secret"))
	if err != nil {
		logger.Errorf("generate refresh token err: %v", err)
		return
	}

	err = db.Orm().Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&models.Token{
			UserId:    accessClaims.UserId,
			Username:  accessClaims.Username,
			JwtId:     accessClaims.ID,
			IssuedAt:  accessClaims.IssuedAt.Unix(),
			ExpiresAt: accessClaims.ExpiresAt.Unix(),
			Status:    models.TokenStatusValid,
			Type:      models.AccessToken,
		}).Error
		if err != nil {
			logger.Errorf("create access token err: %v", err)
			return
		}

		err = tx.Create(&models.Token{
			JwtId:     refreshClaims.ID,
			IssuedAt:  refreshClaims.IssuedAt.Unix(),
			ExpiresAt: refreshClaims.ExpiresAt.Unix(),
			Status:    models.TokenStatusValid,
			Type:      models.RefreshToken,
		}).Error
		if err != nil {
			logger.Errorf("create refresh token err: %v", err)
			return
		}

		return
	})
	if err != nil {
		return
	}

	return &TokenPair{
		AccessToken:  signedAccess,
		RefreshToken: signedRefresh,
		ExpiresAt:    accessClaims.ExpiresAt.Unix(),
	}, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
