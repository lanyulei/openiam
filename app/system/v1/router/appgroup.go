package router

import (
	"openiam/app/system/v1/apis"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
  @Desc :
*/

func AppGroupRouter(g *gin.RouterGroup) {
	router := g.Group("/app-group")
	{
		router.GET("", apis.AppGroupList)
		router.POST("", apis.CreateAppGroup)
		router.PUT("/:id", apis.UpdateAppGroup)
		router.DELETE("/:id", apis.DeleteAppGroup)
	}
}
