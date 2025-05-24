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
	g.POST("/logout", api.Logout, auth.JWTAuthMiddleware())
}
