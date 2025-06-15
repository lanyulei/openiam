package router

import (
	"openops/app/resource/v1/api"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
  @Desc :
*/

func CloudModelsRouter(g *gin.RouterGroup) {
	router := g.Group("/cloud-models")
	{
		router.GET("", api.GetCloudModels)
		router.POST("", api.CreateCloudModel)
		router.PUT("/:id", api.UpdateCloudModel)
		router.DELETE("/:id", api.DeleteCloudModel)
	}
}
