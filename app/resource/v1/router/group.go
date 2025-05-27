package router

import (
	"openops/app/resource/v1/api"

	"github.com/gin-gonic/gin"
)

func MenuRouter(g *gin.RouterGroup) {
	router := g.Group("/model-group")
	{
		router.GET("", api.ModelGroupList)
	}
}
