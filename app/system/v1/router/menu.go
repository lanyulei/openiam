package router

import (
	"openops/app/system/v1/api"

	"github.com/gin-gonic/gin"
)

func MenuRouter(g *gin.RouterGroup) {
	router := g.Group("/menu")
	{
		router.GET("", api.MenuList)
		router.POST("", api.CreateMenu)
		router.PUT("/:id", api.UpdateMenu)
		router.DELETE("/:id", api.DeleteMenu)
		router.GET("/:id", api.MenuDetailByMenuId)
		router.GET("/tree", api.MenuTree)
	}
}
