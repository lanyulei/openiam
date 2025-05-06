package routers

import (
	"openiam/app/system/v1/router"
	"openiam/common/middleware/auth"
	"openiam/common/middleware/permission"

	"github.com/gin-gonic/gin"
)

func RegisterSystemRouter(g *gin.RouterGroup) {
	group := g.Group("/system", auth.JWTAuthMiddleware(), permission.CheckPermMiddleware())

	router.UserRouter(group)             // 用户管理
	router.DepartmentRouter(group)       // 部门管理
	router.MenuRouter(group)             // 菜单管理
	router.RoleRouter(group)             // 角色管理
	router.ApiRouter(group)              // 接口管理
	router.ApiGroupRouter(group)         // 接口分组管理
	router.LoginLogRouter(group)         // 登陆日志
	router.UploadRouter(group)           // 文件上传
	router.SettingsRouter(group)         // 配置管理
	router.AuditRouter(group)            // 审计日志
	router.AppGroupRouter(group)         // 应用分组
	router.AppRouter(group)              // 应用管理
	router.UserGroupRouter(group)        // 用户分组
	router.UserGroupRelatedRouter(group) // 用户分组关联

	// 不要验证的接口
	notVerify := g.Group("/system")
	router.SettingsNotVerifyRouter(notVerify)
}
