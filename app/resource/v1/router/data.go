package router

import (
	"openops/app/resource/v1/api"

	"github.com/gin-gonic/gin"
)

func DataRouter(g *gin.RouterGroup) {
	router := g.Group("/data")
	{
		router.GET("/:id", api.DataList)
		router.GET("/detail/:id", api.DataDetails)
		router.POST("", api.CreateData)
		router.PUT("/:id", api.UpdateData)
		router.DELETE("/batch", api.BatchDeleteData)
	}
}
