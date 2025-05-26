package router

import (
	"openops/app/system/v1/api"

	"github.com/gin-gonic/gin"
)

func UserRouter(g *gin.RouterGroup) {
	router := g.Group("/user")
	{
		router.GET("", api.UserList)
		router.POST("", api.CreateUser)
		router.PUT("/:id", api.UpdateUser)
		router.DELETE("/:id", api.DeleteUser)
		router.GET("/:id", api.UserDetailByUserId)
		router.GET("/details", api.UserDetail)
	}
}
