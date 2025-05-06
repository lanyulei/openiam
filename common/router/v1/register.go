package v1

import (
	"openiam/common/router/v1/routers"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
  @Desc :
*/

func RegisterRouter(g *gin.RouterGroup) {
	routers.RegisterPublicRouter(g) // 公共接口路由，非业务相关路由
	routers.RegisterSystemRouter(g) // 系统管理
}

func MicroServiceRouter(g *gin.Engine) {
}
