package router

import (
	"openiam/app/public/v1/api"
	"openiam/common/middleware/auth"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
*/

func LoginRouter(g *gin.RouterGroup) {
	g.POST("/login", api.Login)
	logoutRouter := g.Group("/logout", auth.JWTAuthMiddleware())
	{
		logoutRouter.POST("", api.Logout)
	}
}
