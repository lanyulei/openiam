package router

import (
	"openops/app/resource/v1/api"

	"github.com/gin-gonic/gin"
)

func LogicHandleRouter(g *gin.RouterGroup) {
	router := g.Group("/logic-handle")
	{
		router.GET("/:id", api.LogicHandleList) // 获取逻辑处理列表
		router.POST("", api.CreateLogicHandle)
		router.PUT("/:id", api.UpdateLogicHandle)
		router.DELETE("/:id", api.DeleteLogicHandle)
	}
}
