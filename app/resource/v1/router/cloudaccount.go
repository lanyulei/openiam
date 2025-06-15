package router

import (
	"openops/app/resource/v1/api"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
  @Desc :
*/

func CloudAccountRouter(g *gin.RouterGroup) {
	router := g.Group("/cloud-account")
	{
		router.GET("", api.CloudAccountList)
		router.POST("", api.CreateCloudAccount)
		router.GET("/:id", api.CloudAccountDetail)
		router.PUT("/:id", api.EditCloudAccount)
		router.DELETE("/:id", api.DeleteCloudAccount)
		router.POST("/check-connect/:id", api.CloudAccountCheckConnect)
		router.POST("/sync-resource", api.SyncCloudResource)
	}
}
