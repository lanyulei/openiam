package router

import (
	"openiam/app/system/v1/apis"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
  @Desc :
*/

func AppRouter(g *gin.RouterGroup) {
	router := g.Group("/app")
	{
		router.GET("", apis.AppList)
		router.GET("/:id", apis.AppListByGroupId)
		router.GET("/list", apis.AppListByGroup)
		router.POST("", apis.CreateApp)
		router.PUT("/:id", apis.UpdateApp)
		router.DELETE("/:id", apis.DeleteApp)
	}
}
