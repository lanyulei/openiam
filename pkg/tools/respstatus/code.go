package respstatus

import "github.com/lanyulei/toolkit/response"

var (
	AuthorizationNullError = response.Response{Code: 30001, Message: "请求头中 Authorization 为空"}
	InvalidTokenError      = response.Response{Code: 30003, Message: "Token 无效"}
	NoPermissionError      = response.Response{Code: 30004, Message: "暂无请求权限"}
	InvalidParamsError     = response.Response{Code: 30005, Message: "请求参数错误"}

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
)
