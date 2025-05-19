package api

import (
	"openiam/app/system/models"
	"openiam/pkg/password"
	"openiam/pkg/tools/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
)

// UserList 用户列表
func UserList(c *gin.Context) {
	var (
		err      error
		userList []*models.User
		result   interface{}
	)

	dbConn := db.Orm().Model(&models.User{})

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
		err  error
		user struct {
			models.User
			Password string `json:"password" binding:"required"`
		}
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

	if user.Password == "" {
		response.Error(c, err, respstatus.PasswordEmptyError)
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

	response.OK(c, user, "")
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

	err = db.Orm().Model(&models.User{}).Where("id = ?", userId).Updates(map[string]interface{}{
		"username": user.Username,
		"nickname": user.Nickname,
		"avatar":   user.Avatar,
		"tel":      user.Tel,
		"email":    user.Email,
		"status":   user.Status,
		"is_admin": user.IsAdmin,
		"remark":   user.Remark,
	}).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateUserError)
		return
	}

	response.OK(c, user, "")
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

// UserDetailByUserId 用户详情
func UserDetailByUserId(c *gin.Context) {
	var (
		err    error
		userId string
		user   models.User
	)

	userId = c.Param("id")

	err = db.Orm().Model(&models.User{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		response.Error(c, err, respstatus.UserDetailError)
		return
	}

	response.OK(c, user, "")
}

// UserDetail 用户详情
func UserDetail(c *gin.Context) {
	var (
		err    error
		userId string
		user   models.User
	)

	userId = c.GetString("userId")

	err = db.Orm().Model(&models.User{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		response.Error(c, err, respstatus.UserDetailError)
		return
	}
	response.OK(c, user, "")
}
