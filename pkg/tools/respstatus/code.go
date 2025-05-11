package respstatus

import "github.com/lanyulei/toolkit/response"

var (
	AuthorizationNullError = response.Response{Code: 30001, Message: "请求头中 Authorization 为空"}
	InvalidTokenError      = response.Response{Code: 30003, Message: "Token 无效"}
	NoPermissionError      = response.Response{Code: 30004, Message: "暂无请求权限"}
	InvalidParamsError     = response.Response{Code: 30005, Message: "请求参数错误"}
)
