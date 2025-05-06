package router

import (
	"openiam/app/system/v1/apis"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
  @Desc :
*/

func UploadRouter(g *gin.RouterGroup) {
	router := g.Group("/upload")
	{
		router.POST("", apis.Upload)
	}
}
