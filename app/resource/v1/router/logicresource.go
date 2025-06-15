package router

import (
	"openops/app/resource/v1/api"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
  @Desc :
*/

func LogicResourceRouter(g *gin.RouterGroup) {
	router := g.Group("/logic-resource")
	{
		router.GET("", api.LogicResourceList)
		router.GET("/:id", api.LogicResourceDetails)
		router.POST("", api.CreateLogicResource)
		router.PUT("/:id", api.UpdateLogicResource)
		router.DELETE("/:id", api.DeleteLogicResource)
	}
}
