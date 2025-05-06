package apis

import (
	"openiam/app/system/models"
	"openiam/pkg/jwtauth"
	"openiam/pkg/tools/respstatus"
	"openiam/server/login"
	"openiam/server/loginlog"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/response"
)

// Login 登陆
func Login(c *gin.Context) {
	var (
		err       error
		token     string
		user      models.UserRequest
		loginUser struct {
			Username string `json:"username"`
			Password string `json:"password"`
			IsLdap   bool   `json:"isLdap"`
		}
	)
	err = c.ShouldBind(&loginUser)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	user, err = login.LoginHandler(loginUser.Username, loginUser.Password, loginUser.IsLdap)
	if err != nil {
		response.Error(c, err, respstatus.LoginError)
		return
	}

	token, err = jwtauth.GenToken(&user)
	if err != nil {
		response.Error(c, err, respstatus.GenerateTokenError)
		return
	}

	go loginlog.Create(c, user.Username, "登陆成功")

	response.OK(c, token, "")
}
