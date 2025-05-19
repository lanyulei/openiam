package api

import (
	"fmt"
	"openiam/app/system/models"
	"openiam/pkg/jwtauth"
	"openiam/pkg/tools/respstatus"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/response"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// Login 登陆
func Login(c *gin.Context) {
	var (
		err    error
		params struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		userInfo models.User
		token    *jwtauth.TokenPair
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

	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(params.Password))
	if err != nil {
		response.Error(c, fmt.Errorf("password error: %v", err), respstatus.CompareHashAndPasswordError)
		return
	}

	token, err = jwtauth.GenerateTokens(userInfo.Id, userInfo.Username)
	if err != nil {
		response.Error(c, err, respstatus.GenerateTokenError)
		return
	}

	c.SetCookie(
		"refresh_token",
		token.RefreshToken,
		int((time.Duration(viper.GetInt("jwt.refreshToken.expires")) * time.Hour).Seconds()), // 过期时间
		"/refresh-token", // 有效路径
		"",               // 域名
		false,            // 仅HTTPS
		true,             // HttpOnly
	)

	response.OK(c, token.AccessToken, "")
}

func RefreshToken(c *gin.Context) {
	response.OK(c, "", "")
}
