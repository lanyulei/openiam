package router

import (
	"openops/app/resource/v1/api"

	"github.com/gin-gonic/gin"
)

func LogicHandleRouter(g *gin.RouterGroup) {
	router := g.Group("/logic-handle")
	{
		router.GET("", api.LogicHandleListById)
		router.POST("", api.CreateLogicHandle)
		router.PUT("/:id", api.UpdateLogicHandle)
		router.DELETE("/:id", api.DeleteLogicHandle)
	}
}
