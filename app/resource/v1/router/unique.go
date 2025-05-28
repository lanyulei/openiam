package router

import (
	"openops/app/resource/v1/api"

	"github.com/gin-gonic/gin"
)

func ModelUniqueRouter(g *gin.RouterGroup) {
	router := g.Group("/model-unique")
	{
		router.GET("/:id", api.ModelUniqueList)
		router.POST("", api.CreateModelUnique)
		router.PUT("/:id", api.UpdateModelUnique)
		router.DELETE("/:id", api.DeleteModelUnique)
	}
}
