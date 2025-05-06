package router

import (
	"openiam/app/public/v1/apis"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
  @Desc :
*/

func RegisterRouter(g *gin.RouterGroup) {
	g.POST("/route/register/check", apis.CheckRegisterRoute)
}
