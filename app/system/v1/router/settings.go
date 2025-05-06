package router

import (
	"openiam/app/system/v1/apis"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
  @Desc :
*/

func SettingsRouter(g *gin.RouterGroup) {
	router := g.Group("/settings")
	{
		router.POST("", apis.UpdateSettings)
	}
}

func SettingsNotVerifyRouter(g *gin.RouterGroup) {
	router := g.Group("/settings")
	{
		router.GET("", apis.GetSettings)
	}
}
