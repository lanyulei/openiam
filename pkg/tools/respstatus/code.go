package respstatus

import "github.com/lanyulei/toolkit/response"

var (
	AuthorizationNullError   = response.Response{Code: 30001, Message: "请求头中 Authorization 为空"}
	AuthorizationFormatError = response.Response{Code: 30002, Message: "请求头中 Authorization 格式有误"}
	InvalidTokenError        = response.Response{Code: 30003, Message: "Token 无效"}
	NoPermissionError        = response.Response{Code: 30004, Message: "暂无请求权限"}

	InvalidParameterError  = response.Response{Code: 40001, Message: "无效参数"}
	QueryUserError         = response.Response{Code: 40002, Message: "查询用户失败"}
	UserExistError         = response.Response{Code: 40004, Message: "用户名已存在"}
	IncorrectPasswordError = response.Response{Code: 40005, Message: "密码不正确"}
	GenerateTokenError     = response.Response{Code: 40006, Message: "生成 Token 失败"}
	CreateUserError        = response.Response{Code: 40007, Message: "创建用户失败"}
	EncryptPasswordError   = response.Response{Code: 40008, Message: "加密密码失败"}
	GetUserInfoError       = response.Response{Code: 40009, Message: "获取用户详情失败"}
	GetMenuError           = response.Response{Code: 40010, Message: "查询菜单数据失败"}
	SaveMenuError          = response.Response{Code: 40011, Message: "保存菜单数据失败"}
	SubmenuExistsError     = response.Response{Code: 40012, Message: "当前菜单存在子节点，无法直接删除"}
	DeleteMenuError        = response.Response{Code: 40013, Message: "删除菜单失败"}
	GetMenuButtonError     = response.Response{Code: 40014, Message: "查询菜单按钮数据失败"}
	UserListError          = response.Response{Code: 40015, Message: "查询用户列表失败"}
	DeleteUserError        = response.Response{Code: 40016, Message: "删除用户失败"}
	UpdateUserError        = response.Response{Code: 40017, Message: "更新用户失败"}
	RoleListError          = response.Response{Code: 40019, Message: "查询角色列表失败"}
	SaveRoleError          = response.Response{Code: 40020, Message: "保存角色失败"}
	RoleUsedError          = response.Response{Code: 40021, Message: "角色被其他用户关联，无法删除"}
	DeleteRoleError        = response.Response{Code: 40022, Message: "角色删除失败"}
	ApiListError           = response.Response{Code: 40023, Message: "查询API接口列表失败"}
	SaveApiError           = response.Response{Code: 40024, Message: "保存API接口失败"}
	DeleteApiError         = response.Response{Code: 40025, Message: "API接口删除失败"}
	GetApiMenuError        = response.Response{Code: 40026, Message: "获取API接口对应的菜单或按钮失败"}
	ApiUsedError           = response.Response{Code: 40027, Message: "API接口被其他角色关联，无法删除"}
	SaveApiGroupError      = response.Response{Code: 40028, Message: "保存API分组失败"}
	ApiGroupListError      = response.Response{Code: 40029, Message: "获取API分组列表失败"}
	ApiGroupExistError     = response.Response{Code: 40030, Message: "API分组已存在"}
	RoleExistError         = response.Response{Code: 40031, Message: "角色已存在"}
	GetApiError            = response.Response{Code: 40032, Message: "获取API失败"}
	ApiGroupUsedError      = response.Response{Code: 40033, Message: "API分组被使用，无法删除"}
	DeleteApiGroupError    = response.Response{Code: 40034, Message: "删除API分组失败"}
	GetRoleMenuError       = response.Response{Code: 40035, Message: "查询角色对应的菜单列表失败"}
	CreateRoleMenuError    = response.Response{Code: 40036, Message: "角色关联菜单权限失败"}
	DeleteRoleMenuError    = response.Response{Code: 40037, Message: "删除角色关联菜单权限失败"}
	GetRolePermissionError = response.Response{Code: 40038, Message: "查询角色对应的菜单权限失败"}
	GetRoleButtonError     = response.Response{Code: 40039, Message: "查询角色对应的菜单按钮权限失败"}
	GetMenuParentError     = response.Response{Code: 40040, Message: "查询所有父级别菜单失败"}
	GetMenuApiError        = response.Response{Code: 40041, Message: "查询菜单绑定的API失败"}
	MenuBindApiError       = response.Response{Code: 40042, Message: "菜单绑定API失败"}
	MenuUnBindApiError     = response.Response{Code: 40043, Message: "菜单结束绑定API失败"}
	CreateUserRoleError    = response.Response{Code: 40044, Message: "创建用户与角色的关联到Casbin中失败"}
	DeleteUserRoleError    = response.Response{Code: 40045, Message: "删除用户与角色的关联到Casbin中失败"}
	GetRoleError           = response.Response{Code: 40046, Message: "获取角色信息失败"}
	UserRoleError          = response.Response{Code: 40047, Message: "当前用户关联了其他角色，无法直接删除"}
	RoleBindApiError       = response.Response{Code: 40048, Message: "角色绑定接口权限失败"}
	RoleUnBindApiError     = response.Response{Code: 40049, Message: "角色解除接口权限失败"}
	RoleBindMenuError      = response.Response{Code: 40050, Message: "角色存在与菜单的绑定关联，无法删除"}
	LoginLogListError      = response.Response{Code: 40051, Message: "查询登陆日志失败"}
	DeleteLoginLogError    = response.Response{Code: 40052, Message: "删除登陆日志失败"}

	GetUserError = response.Response{Code: 40092, Message: "获取用户失败"}

	LoginLogInfoError          = response.Response{Code: 40110, Message: "获取当前登录用户的最近一次登录详情失败"}
	UpdatePermissionCacheError = response.Response{Code: 40111, Message: "更新角色权限缓存失败"}
	GetAppGroupError           = response.Response{Code: 40112, Message: "获取应用分组失败"}
	CreateAppGroupError        = response.Response{Code: 40113, Message: "创建应用分组失败"}
	UpdateAppGroupError        = response.Response{Code: 40114, Message: "更新应用分组失败"}
	DeleteAppGroupError        = response.Response{Code: 40115, Message: "删除应用分组失败"}
	AppGroupExistError         = response.Response{Code: 40116, Message: "应用分组已存在"}
	AppGroupHasAppError        = response.Response{Code: 40117, Message: "应用分组下存在应用"}
	GetAppError                = response.Response{Code: 40118, Message: "获取应用失败"}
	CreateAppError             = response.Response{Code: 40119, Message: "创建应用失败"}
	UpdateAppError             = response.Response{Code: 40120, Message: "更新应用失败"}
	DeleteAppError             = response.Response{Code: 40121, Message: "删除应用失败"}
	AppExistError              = response.Response{Code: 40122, Message: "应用已存在"}
	AppHasMenuError            = response.Response{Code: 40123, Message: "应用下存在菜单"}
	DecodePasswordError        = response.Response{Code: 40124, Message: "解码密码失败"}

	GetDepartmentListError      = response.Response{Code: 40200, Message: "获取部门列表失败"}
	CreateDepartmentError       = response.Response{Code: 40201, Message: "创建部门失败"}
	UpdateDepartmentError       = response.Response{Code: 40202, Message: "更新部门失败"}
	DeleteDepartmentError       = response.Response{Code: 40203, Message: "删除部门失败"}
	GetDepartmentError          = response.Response{Code: 40204, Message: "获取部门信息失败"}
	CreateUserGroupError        = response.Response{Code: 40205, Message: "创建用户组失败"}
	UpdateUserGroupError        = response.Response{Code: 40206, Message: "更新用户组失败"}
	UserGroupRelatedExistError  = response.Response{Code: 40207, Message: "用户组下存在用户"}
	DeleteUserGroupError        = response.Response{Code: 40208, Message: "删除用户组失败"}
	UserGroupListError          = response.Response{Code: 40209, Message: "获取用户组列表失败"}
	UserIdsError                = response.Response{Code: 40210, Message: "获取用户 ID 列表失败"}
	CreateUserGroupRelatedError = response.Response{Code: 40211, Message: "创建用户组关联失败"}
	DeleteUserGroupRelatedError = response.Response{Code: 40212, Message: "删除用户组关联失败"}
	GetApiGroupError            = response.Response{Code: 40213, Message: "获取API分组失败"}
	LoginError                  = response.Response{Code: 40214, Message: "登录失败"}
	GetRoleUserError            = response.Response{Code: 40215, Message: "获取角色下的用户失败"}
	GetRoleApiError             = response.Response{Code: 40216, Message: "获取角色下的接口失败"}

	UploadFileError         = response.Response{Code: 40400, Message: "上传文件失败"}
	GetSettingsError        = response.Response{Code: 40401, Message: "获取系统配置失败"}
	UpdateSettingsError     = response.Response{Code: 40402, Message: "更新系统配置失败"}
	CheckRegisterRouteError = response.Response{Code: 40405, Message: "检查注册路由失败"}

	GetAuditLogError    = response.Response{Code: 40600, Message: "获取审计日志失败"}
	DeleteAuditLogError = response.Response{Code: 40601, Message: "删除审计日志失败"}

	RequestForwardError = response.Response{Code: 41002, Message: "请求转发失败"}
)
