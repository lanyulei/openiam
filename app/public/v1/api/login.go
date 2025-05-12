package api

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/response"
	"golang.org/x/crypto/bcrypt"
	"openiam/app/system/models"
	"openiam/pkg/tools/respstatus"
)

// Login 登陆
func Login(c *gin.Context) {
	var (
		err    error
		params struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		decodedPassword []byte
		userInfo        models.User
	)

	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// 查询用户信息
	err = db.Orm().Model(&models.User{}).Where("username = ?", params.Username).First(&userInfo).Error
	if err != nil {
		response.Error(c, err, respstatus.GetUserError)
		return
	}

	decodedPassword, err = base64.StdEncoding.DecodeString(params.Password)
	if err != nil {
		response.Error(c, fmt.Errorf("decode password error: %v", err), respstatus.DecodedPasswordError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), decodedPassword)
	if err != nil {
		response.Error(c, fmt.Errorf("password error: %v", err), respstatus.CompareHashAndPasswordError)
		return
	}

	response.OK(c, "", "")
}

func RefreshToken(c *gin.Context) {
	response.OK(c, "", "")
}
