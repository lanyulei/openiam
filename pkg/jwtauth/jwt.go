package jwtauth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"time"
)

type Claims struct {
	UserId   int    `json:"user_id"`
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

func GenerateTokens(userId int, username string) (result *TokenPair, err error) {
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
		return nil, err
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
		return nil, err
	}

	// 存储Refresh Token到数据库（需要实现）
	//if err := storeRefreshToken(userId, refreshClaims.ID, refreshClaims.ExpiresAt.Time); err != nil {
	//	return nil, err
	//}

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
