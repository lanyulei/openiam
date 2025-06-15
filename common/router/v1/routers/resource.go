package routers

import (
	"openops/app/resource/v1/router"
	"openops/common/middleware/auth"

	"github.com/gin-gonic/gin"
)

func RegisterResourceRouter(g *gin.RouterGroup) {
	group := g.Group("/resource", auth.JWTAuthMiddleware())
	router.ModelGroupRouter(group) // 模型分组
	router.ModelRouter(group)      // 模型
	router.FieldGroupRouter(group) // 字段分组
	router.FieldRouter(group)      // 字段
	router.DataRouter(group)       // 数据

	router.LogicResourceRouter(group) // 逻辑资源
	router.LogicHandleRouter(group)   // 逻辑处理
	router.CloudRegionRouter(group)   // 云区域
	router.CloudModelsRouter(group)   // 云模型
	router.CloudAccountRouter(group)  // 云账号

}
