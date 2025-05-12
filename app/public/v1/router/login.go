package router

import (
	"openiam/app/public/v1/api"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
*/

func LoginRouter(g *gin.RouterGroup) {
	g.POST("/login", api.Login)
	g.POST("/refresh-token", api.RefreshToken)
}
