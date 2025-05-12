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
		router.PUT("/:id", api.UpdateUser)
		router.DELETE("/:id", api.DeleteUser)
	}
}
