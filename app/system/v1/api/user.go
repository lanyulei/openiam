package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
	"openiam/app/system/models"
	"openiam/pkg/password"
	"openiam/pkg/tools/respstatus"
)

// UserList 用户列表
func UserList(c *gin.Context) {
	var (
		err      error
		userList []*struct {
			models.User
			DepartmentName string `json:"department_name"`
		}
		result interface{}
	)

	dbConn := db.Orm().Model(&models.User{}).
		Select("system_user.*, system_department.name as department_name").
		Joins("left join system_department on system_department.id = system_user.department")

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: dbConn,
	}, &userList)
	if err != nil {
		response.Error(c, err, respstatus.UserListError)
		return
	}

	response.OK(c, result, "")
}

// CreateUser 创建用户
func CreateUser(c *gin.Context) {
	var (
		err   error
		user  models.User
		count int64
	)

	err = c.ShouldBindJSON(&user)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// 用户名唯一性校验
	err = db.Orm().Model(&models.User{}).Where("username = ?", user.Username).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetUserError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.UsernameExistError)
		return
	}

	user.Password, err = password.EncryptionPassword(user.Password)
	if err != nil {
		response.Error(c, err, respstatus.EncryptionPasswordError)
		return
	}

	err = db.Orm().Create(&user).Error
	if err != nil {
		response.Error(c, err, respstatus.CreateUserError)
		return
	}

	response.OK(c, "", "")
}

// UpdateUser 更新用户
func UpdateUser(c *gin.Context) {
	var (
		err    error
		user   models.User
		count  int64
		userId string
	)

	userId = c.Param("id")

	err = c.ShouldBindJSON(&user)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// 用户名唯一性校验，排除自己
	err = db.Orm().Model(&models.User{}).Where("username = ? and id != ?", user.Username, userId).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetUserError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.UsernameExistError)
		return
	}

	err = db.Orm().Model(&models.User{}).Where("id = ?", user.Id).Updates(&user).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateUserError)
		return
	}

	response.OK(c, "", "")
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	var (
		err    error
		userId string
	)

	userId = c.Param("id")

	err = db.Orm().Model(&models.User{}).Where("id = ?", userId).Delete(&models.User{}).Error
	if err != nil {
		response.Error(c, err, respstatus.GetUserError)
		return
	}

	response.OK(c, "", "")
}
