package router

import (
	"openiam/app/system/v1/apis"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
*/

func UserRouter(g *gin.RouterGroup) {
	router := g.Group("/user")
	{
		router.GET("", apis.UserList)
		router.GET("/list", apis.UserListById)
		router.GET("/group", apis.UserListByGroupId)
		router.GET("/info", apis.UserInfo)
		router.PUT("/init-password/:id", apis.InitPassword)
		router.PUT("/update-password", apis.UpdatePassword)
		router.GET("/info/:id", apis.UserInfoById)
		router.PUT("/info", apis.UpdateUserInfo)
		router.GET("/details/:username", apis.UserInfoByUsername)
		router.POST("", apis.CreateUser)
		router.PUT("/:id", apis.UpdateUser)
		router.DELETE("/:id", apis.DeleteUser)
		router.PUT("/avatar/:id", apis.UpdateUserAvatar)
	}
}
