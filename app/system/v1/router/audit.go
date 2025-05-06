package router

import (
	"openiam/app/system/v1/apis"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
  @Desc :
*/

func AuditRouter(g *gin.RouterGroup) {
	router := g.Group("/audit")
	{
		router.GET("", apis.AuditList)
		router.DELETE("/:id", apis.DeleteAudit)
		router.GET("/:id", apis.GetAuditInfo)
	}
}
