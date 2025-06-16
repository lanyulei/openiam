package router

import (
	"github.com/gin-gonic/gin"
	"openops/app/resource/v1/api"
)

/*
  @Author : lanyulei
  @Desc :
*/

func PluginRouter(g *gin.RouterGroup) {
	router := g.Group("/plugin")
	{
		router.GET("", api.PluginList)
	}
}
