package router

import (
	"github.com/gin-gonic/gin"
	"openiam/app/system/v1/api"
)

func UserRouter(g *gin.RouterGroup) {
	router := g.Group("/user")
	{
		router.GET("", api.UserList)
		router.POST("", api.CreateUser)
		router.PUT("", api.UpdateUser)
		router.DELETE("", api.DeleteUser)
	}
}
