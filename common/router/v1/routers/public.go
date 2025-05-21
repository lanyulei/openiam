package routers

import (
	"openiam/app/public/v1/router"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRouter(g *gin.RouterGroup) {
	router.LoginRouter(g)
	router.MenuRouter(g)
}
