package respstatus

import "github.com/lanyulei/toolkit/response"

var (
	AuthorizationNullError = response.Response{Code: 30001, Message: "请求头中 Authorization 为空"}
	InvalidTokenError      = response.Response{Code: 30003, Message: "Token 无效"}
	NoPermissionError      = response.Response{Code: 30004, Message: "暂无请求权限"}
	InvalidParamsError     = response.Response{Code: 30005, Message: "请求参数错误"}
	TokenNotFoundError     = response.Response{Code: 30008, Message: "令牌不存在"}
	UnmarshalError         = response.Response{Code: 30009, Message: "反序列化错误"}

	UserListError               = response.Response{Code: 40001, Message: "获取用户列表失败"}
	GetUserError                = response.Response{Code: 40002, Message: "获取用户失败"}
	UsernameExistError          = response.Response{Code: 40003, Message: "用户名已存在"}
	EncryptionPasswordError     = response.Response{Code: 40004, Message: "加密密码失败"}
	CreateUserError             = response.Response{Code: 40005, Message: "创建用户失败"}
	UpdateUserError             = response.Response{Code: 40006, Message: "更新用户失败"}
	CompareHashAndPasswordError = response.Response{Code: 40008, Message: "密码错误"}
	GenerateTokenError          = response.Response{Code: 40009, Message: "生成 token 失败"}
	PasswordEmptyError          = response.Response{Code: 40010, Message: "密码不能为空"}
	UserDetailError             = response.Response{Code: 40011, Message: "获取用户详情失败"}
	UserNotFoundError           = response.Response{Code: 40012, Message: "用户不存在"}
	InvalidLoginTypeError       = response.Response{Code: 40013, Message: "无效的登录类型"}

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
	ModelListError           = response.Response{Code: 40208, Message: "获取模型列表失败"}
	GetModelError            = response.Response{Code: 40209, Message: "获取模型失败"}
	ModelNameExistError      = response.Response{Code: 40210, Message: "模型名称已存在"}
	CreateModelError         = response.Response{Code: 40211, Message: "创建模型失败"}
	UpdateModelError         = response.Response{Code: 40212, Message: "更新模型失败"}
	DeleteModelError         = response.Response{Code: 40213, Message: "删除模型失败"}
	GetModelDataError        = response.Response{Code: 40218, Message: "获取模型数据失败"}
	ModelHasDataError        = response.Response{Code: 40219, Message: "模型下有数据，无法删除"}
	FieldGroupListError      = response.Response{Code: 40220, Message: "获取模型字段组列表失败"}
	GetFieldGroupError       = response.Response{Code: 40221, Message: "获取模型字段组失败"}
	FieldGroupNameExistError = response.Response{Code: 40222, Message: "模型字段组名称已存在"}
	CreateFieldGroupError    = response.Response{Code: 40223, Message: "创建模型字段组失败"}
	UpdateFieldGroupError    = response.Response{Code: 40224, Message: "更新模型字段组失败"}
	DeleteFieldGroupError    = response.Response{Code: 40225, Message: "删除模型字段组失败"}
	FieldGroupHasBindError   = response.Response{Code: 40226, Message: "模型字段组下有绑定，无法删除"}
	FieldListError           = response.Response{Code: 40227, Message: "获取模型字段列表失败"}
	FieldExistError          = response.Response{Code: 40228, Message: "模型字段名称已存在"}
	GetFieldError            = response.Response{Code: 40229, Message: "获取模型字段失败"}
	CreateFieldError         = response.Response{Code: 40230, Message: "创建模型字段失败"}
	UpdateFieldError         = response.Response{Code: 40231, Message: "更新模型字段失败"}
	DeleteFieldError         = response.Response{Code: 40232, Message: "删除模型字段失败"}
	GetModelRelationError    = response.Response{Code: 40233, Message: "获取模型关联失败"}
	CreateModelRelationError = response.Response{Code: 40235, Message: "创建模型关联失败"}
	ModelRelationExistError  = response.Response{Code: 40237, Message: "模型关联已存在"}
	UpdateModelRelationError = response.Response{Code: 40238, Message: "更新模型关联失败"}
	DeleteModelRelationError = response.Response{Code: 40239, Message: "删除模型关联失败"}
	GetModelUniqueError      = response.Response{Code: 40240, Message: "获取模型唯一约束失败"}
	CreateModelUniqueError   = response.Response{Code: 40241, Message: "创建模型唯一约束失败"}
	ModelUniqueExistError    = response.Response{Code: 40242, Message: "模型唯一约束已存在"}
	UpdateModelUniqueError   = response.Response{Code: 40243, Message: "更新模型唯一约束失败"}
	DeleteModelUniqueError   = response.Response{Code: 40244, Message: "删除模型唯一约束失败"}
	DataListError            = response.Response{Code: 40245, Message: "获取模型数据列表失败"}
	FieldKeyExistError       = response.Response{Code: 40246, Message: "模型字段键已存在"}
	VerifyDataError          = response.Response{Code: 40247, Message: "验证数据失败"}
	CreateDataError          = response.Response{Code: 40248, Message: "创建数据失败"}
	UpdateDataError          = response.Response{Code: 40249, Message: "更新数据失败"}
	BuiltInFieldError        = response.Response{Code: 40250, Message: "不能是内置字段"}
	BatchDeleteDataError     = response.Response{Code: 40251, Message: "批量删除数据失败"}
	GetDataDetailsError      = response.Response{Code: 40252, Message: "获取数据详情失败"}
)
