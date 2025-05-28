package router

import (
	"openops/app/resource/v1/api"

	"github.com/gin-gonic/gin"
)

func ModelRouter(g *gin.RouterGroup) {
	router := g.Group("/model")
	{
		router.GET("", api.ModelGroupList)
		router.POST("", api.CreateModelGroup)
		router.PUT("/:id", api.UpdateModelGroup)
		router.DELETE("/:id", api.DeleteModelGroup)
	}
}
