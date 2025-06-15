package router

import (
	"openops/app/resource/v1/api"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
  @Desc :
*/

func CloudRegionRouter(g *gin.RouterGroup) {
	router := g.Group("/cloud-region")
	{
		router.GET("", api.GetRegions)
		router.POST("", api.CreateRegion)
		router.PUT("/:id", api.UpdateRegion)
		router.DELETE("/:id", api.DeleteRegion)
	}
}
