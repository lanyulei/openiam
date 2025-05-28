package router

import (
	"openops/app/resource/v1/api"

	"github.com/gin-gonic/gin"
)

func FieldGroupRouter(g *gin.RouterGroup) {
	router := g.Group("/field-group")
	{
		router.GET("", api.FieldGroupList)
	}
}
