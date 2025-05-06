package router

import (
	"openiam/app/system/v1/apis"

	"github.com/gin-gonic/gin"
)

func ApiRouter(g *gin.RouterGroup) {
	router := g.Group("/api")
	{
		router.GET("", apis.ApiList)
		router.POST("", apis.SaveApi)
		router.POST("/batch", apis.BatchCreateApi)
		router.DELETE("/:id", apis.DeleteApi)
		router.PUT("/no-forensics", apis.UpdateApiNoForensics)
	}
}
