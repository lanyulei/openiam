package router

import (
	"openiam/app/system/v1/apis"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
*/

func UserGroupRouter(g *gin.RouterGroup) {
	router := g.Group("/user-group")
	{
		router.GET("", apis.UserGroupList)
		router.GET("/list", apis.UserGroups)
		router.POST("", apis.CreateUserGroup)
		router.PUT("/:id", apis.UpdateUserGroup)
		router.DELETE("/:id", apis.DeleteUserGroup)
	}
}
