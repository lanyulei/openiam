package router

import (
	"openiam/app/public/v1/api"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
*/

func MenuRouter(g *gin.RouterGroup) {
	router := g.Group("/menu")
	{
		router.GET("/constant", api.GetConstantRoutes)
	}
}
