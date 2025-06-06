package router

import (
	"openops/app/resource/v1/api"

	"github.com/gin-gonic/gin"
)

func ModelRouter(g *gin.RouterGroup) {
	router := g.Group("/model")
	{
		router.GET("/list", api.GetModels)
		router.GET("", api.ModelList)
		router.GET("/:id", api.GetModel)
		router.POST("", api.CreateModel)
		router.PUT("/:id", api.UpdateModel)
		router.DELETE("/:id", api.DeleteModel)
	}
}
