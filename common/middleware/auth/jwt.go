package auth

import (
	"openops/app/system/models"
	"openops/pkg/jwtauth"
	"openops/pkg/respstatus"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/response"
	"github.com/spf13/viper"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			user  models.User
			token models.Token
		)

		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			response.Error(c, nil, respstatus.AuthorizationNullError)
			c.Abort()
			return
		}

		mc, err := jwtauth.ParseToken(strings.TrimPrefix(authorization, "Bearer "), viper.GetString("jwt.accessToken.secret"), jwtauth.AccessClaim)
		if err != nil {
			response.Error(c, err, respstatus.InvalidTokenError)
			c.Abort()
			return
		}

		// 检查token是否在数据库中存在
		err = db.Orm().Model(&models.Token{}).Where("jwt_id = ?", mc.(*jwtauth.Claims).ID).First(&token).Error
		if err != nil {
			response.Error(c, err, respstatus.TokenNotFoundError)
			c.Abort()
			return
		}

		// 检查 token 是否有效
		if token.Status != models.TokenStatusValid {
			response.Error(c, nil, respstatus.InvalidTokenError)
			c.Abort()
			return
		}

		// 查询当前用户
		err = db.Orm().Model(&models.User{}).Where("id = ?", mc.(*jwtauth.Claims).UserId).First(&user).Error
		if err != nil {
			response.Error(c, err, respstatus.UserNotFoundError)
			c.Abort()
			return
		}

		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(MiddlewareUsername, mc.(*jwtauth.Claims).Username)
		c.Set(MiddlewareUserId, mc.(*jwtauth.Claims).UserId)
		c.Set(MiddlewareIsAdmin, user.IsAdmin)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
