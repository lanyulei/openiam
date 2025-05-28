package router

import (
	"openops/app/resource/v1/api"

	"github.com/gin-gonic/gin"
)

func ModelRelationRouter(g *gin.RouterGroup) {
	router := g.Group("/model-relation")
	{
		router.GET("/:sourceModelId", api.ModelRelationBySourceModelIdList)
		router.POST("", api.CreateModelRelation)
		router.PUT("/:id", api.UpdateModelRelation)
		router.DELETE("/:id", api.DeleteModelRelation)
	}
}
