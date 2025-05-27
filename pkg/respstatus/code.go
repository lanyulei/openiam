package respstatus

import "github.com/lanyulei/toolkit/response"

var (
	AuthorizationNullError = response.Response{Code: 30001, Message: "请求头中 Authorization 为空"}
	InvalidTokenError      = response.Response{Code: 30003, Message: "Token 无效"}
	NoPermissionError      = response.Response{Code: 30004, Message: "暂无请求权限"}
	InvalidParamsError     = response.Response{Code: 30005, Message: "请求参数错误"}
	GetRefreshTokenError   = response.Response{Code: 30006, Message: "获取刷新令牌失败"}
	ParseRefreshTokenError = response.Response{Code: 30007, Message: "解析刷新令牌失败"}
	TokenNotFoundError     = response.Response{Code: 30008, Message: "令牌不存在"}

	UserListError               = response.Response{Code: 40001, Message: "获取用户列表失败"}
	GetUserError                = response.Response{Code: 40002, Message: "获取用户失败"}
	UsernameExistError          = response.Response{Code: 40003, Message: "用户名已存在"}
	EncryptionPasswordError     = response.Response{Code: 40004, Message: "加密密码失败"}
	CreateUserError             = response.Response{Code: 40005, Message: "创建用户失败"}
	UpdateUserError             = response.Response{Code: 40006, Message: "更新用户失败"}
	DecodedPasswordError        = response.Response{Code: 40007, Message: "解密密码失败"}
	CompareHashAndPasswordError = response.Response{Code: 40008, Message: "密码错误"}
	GenerateTokenError          = response.Response{Code: 40009, Message: "生成 token 失败"}
	PasswordEmptyError          = response.Response{Code: 40010, Message: "密码不能为空"}
	UserDetailError             = response.Response{Code: 40011, Message: "获取用户详情失败"}
	UserNotFoundError           = response.Response{Code: 40012, Message: "用户不存在"}

	GetMenuError           = response.Response{Code: 40101, Message: "获取菜单失败"}
	CreateMenuError        = response.Response{Code: 40102, Message: "创建菜单失败"}
	GetMenuDetailsError    = response.Response{Code: 40103, Message: "获取菜单详情失败"}
	UpdateMenuError        = response.Response{Code: 40104, Message: "更新菜单失败"}
	DeleteMenuError        = response.Response{Code: 40105, Message: "删除菜单失败"}
	GetMenuListError       = response.Response{Code: 40106, Message: "获取菜单列表失败"}
	PathAlreadyExistsError = response.Response{Code: 40107, Message: "路径已存在"}
	GetMenuTreeError       = response.Response{Code: 40108, Message: "获取菜单树失败"}

	GetModelGroupError       = response.Response{Code: 40201, Message: "获取模型组失败"}
	ModelGroupNameExistError = response.Response{Code: 40202, Message: "模型组名称已存在"}
	CreateModelGroupError    = response.Response{Code: 40203, Message: "创建模型组失败"}
	UpdateModelGroupError    = response.Response{Code: 40204, Message: "更新模型组失败"}
	DeleteModelGroupError    = response.Response{Code: 40205, Message: "删除模型组失败"}
	ModelGroupHasModelError  = response.Response{Code: 40206, Message: "模型组下有模型，无法删除"}
)
