package auth

import (
	"openiam/pkg/jwtauth"
	"openiam/pkg/respstatus"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/response"
	"github.com/spf13/viper"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			response.Error(c, nil, respstatus.AuthorizationNullError)
			c.Abort()
			return
		}

		mc, err := jwtauth.ParseToken(strings.TrimPrefix(authorization, "Bearer "), viper.GetString("jwt.accessToken.secret"))
		if err != nil {
			response.Error(c, err, respstatus.InvalidTokenError)
			c.Abort()
			return
		}

		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Set("userId", mc.UserId)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
