package router

import (
	"openiam/app/system/v1/api"

	"github.com/gin-gonic/gin"
)

func MenuRouter(g *gin.RouterGroup) {
	router := g.Group("/menu")
	{
		router.GET("", api.MenuList)
	}
}
