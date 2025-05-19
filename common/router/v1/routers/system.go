package routers

import (
	"openiam/app/system/v1/router"
	"openiam/common/middleware/auth"

	"github.com/gin-gonic/gin"
)

func RegisterSystemRouter(g *gin.RouterGroup) {
	group := g.Group("/system", auth.JWTAuthMiddleware())
	router.UserRouter(group) // 用户管理
}
