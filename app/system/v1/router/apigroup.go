package router

import (
	"openiam/app/system/v1/apis"

	"github.com/gin-gonic/gin"
)

func ApiGroupRouter(g *gin.RouterGroup) {
	router := g.Group("/api-group")
	{
		router.GET("", apis.ApiGroupList)
		router.POST("", apis.SaveApiGroup)
		router.DELETE("/:id", apis.DeleteApiGroup)
	}
}
