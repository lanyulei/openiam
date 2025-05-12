package jwtauth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/logger"
	"github.com/spf13/viper"
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
}

func GenerateTokens(userId string, username string) (result *TokenPair, err error) {
	var (
		now                         = time.Now()
		jti                         = uuid.New().String()
		issuer                      = viper.GetString("jwt.issuer")
		signedAccess, signedRefresh string
	)

	signedAccess, err = GenerateAccessTokens(jti, userId, username, issuer, now)
	if err != nil {
		return
	}

	signedRefresh, err = GenerateRefreshTokens(jti, issuer, now)
	if err != nil {
		return
	}

	return &TokenPair{
		AccessToken:  signedAccess,
		RefreshToken: signedRefresh,
	}, nil
}

// GenerateAccessTokens 生成访问令牌
func GenerateAccessTokens(jti, userId, username, issuer string, now time.Time) (result string, err error) {
	var (
		accessToken *jwt.Token
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
	result, err = accessToken.SignedString([]byte(viper.GetString("jwt.accessToken.secret")))
	if err != nil {
		logger.Errorf("jwt sign access token err: %v", err)
		return
	}

	err = db.Orm().Create(&models.Token{
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

	return
}

// GenerateRefreshTokens 生成刷新令牌
func GenerateRefreshTokens(jti, issuer string, now time.Time) (result string, err error) {
	// Refresh Token（不包含用户敏感信息）
	refreshClaims := &RefreshClaims{
		jti: jti,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(viper.GetInt("jwt.refreshToken.expires")) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    issuer,
			ID:        uuid.New().String(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	result, err = refreshToken.SignedString([]byte(viper.GetString("jwt.accessToken.secret")))
	if err != nil {
		logger.Errorf("generate refresh token err: %v", err)
		return
	}

	err = db.Orm().Create(&models.Token{
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
