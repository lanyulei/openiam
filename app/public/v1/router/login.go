package router

import (
	"openiam/app/public/v1/apis"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
*/

func LoginRouter(g *gin.RouterGroup) {
	g.POST("/login", apis.Login)
}
