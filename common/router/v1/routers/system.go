package routers

import (
	"openops/app/system/v1/router"
	"openops/common/middleware/auth"

	"github.com/gin-gonic/gin"
)

func RegisterSystemRouter(g *gin.RouterGroup) {
	group := g.Group("/system", auth.JWTAuthMiddleware())
	router.UserRouter(group) // 用户管理
	// router.MenuRouter(group) // 菜单管理
}
