package router

import (
	"openiam/app/system/v1/apis"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
*/

func MenuRouter(g *gin.RouterGroup) {
	router := g.Group("/menu")
	{
		router.GET("/tree", apis.MenuTree)
		router.GET("/user-tree", apis.UserMenuTree)
		router.POST("", apis.SaveMenu)
		router.DELETE("/:id", apis.DeleteMenu)
		router.DELETE("/batch", apis.BatchDeleteMenu)
		router.GET("/button/:id", apis.MenuButton)
		router.POST("/api/:id", apis.MenuBindApi)
		router.GET("/api/:id", apis.MenuApis)
		router.GET("/api-list/:id", apis.MenuApiList)
	}
}
