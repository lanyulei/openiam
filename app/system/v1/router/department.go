package router

import (
	"openiam/app/system/v1/apis"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
  @Desc :
*/

func DepartmentRouter(g *gin.RouterGroup) {
	router := g.Group("/department")
	{
		router.GET("", apis.DepartmentList)
		router.GET("/tree", apis.DepartmentTree)
		router.POST("", apis.CreateDepartment)
		router.PUT("/:id", apis.UpdateDepartment)
		router.DELETE("/:id", apis.DeleteDepartment)
	}
}
