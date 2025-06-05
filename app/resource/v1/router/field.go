package router

import (
	"openops/app/resource/v1/api"

	"github.com/gin-gonic/gin"
)

func FieldRouter(g *gin.RouterGroup) {
	router := g.Group("/field")
	{
		router.GET("/list/:id", api.GetFieldsAndGroups)
		router.GET("", api.FieldList)
		router.POST("", api.CreateField)
		router.PUT("/:id", api.UpdateField)
		router.DELETE("/:id", api.DeleteField)
	}
}
