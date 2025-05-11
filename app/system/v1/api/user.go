package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/response"
)

// UserList 用户列表
func UserList(c *gin.Context) {
	response.OK(c, "", "")
}

// CreateUser 创建用户
func CreateUser(c *gin.Context) {
	response.OK(c, "", "")
}

// UpdateUser 更新用户
func UpdateUser(c *gin.Context) {
	response.OK(c, "", "")
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	response.OK(c, "", "")
}
