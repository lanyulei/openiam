package router

import (
	"openiam/app/system/v1/apis"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
*/

func UserGroupRelatedRouter(g *gin.RouterGroup) {
	router := g.Group("/user-group-related")
	{
		router.POST("", apis.CreateUserGroupRelated)
		router.DELETE("", apis.DeleteUserGroupRelated)
	}
}
