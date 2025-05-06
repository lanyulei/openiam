package loginlog

import (
	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/logger"
	"github.com/mssola/user_agent"
	"openiam/app/system/models"
	"openiam/pkg/tools"
)

func Create(c *gin.Context, username, status string) {
	ua := user_agent.New(c.Request.UserAgent())
	browserName, browserVersion := ua.Browser()

	loginLog := models.LoginLog{
		Username: username,
		Status:   status,
		IP:       tools.GetClientIP(c),
		Browser:  browserName + " " + browserVersion,
		System:   ua.OS(),
		Remark:   c.Request.UserAgent(),
	}

	err := db.Orm().Create(&loginLog).Error
	if err != nil {
		logger.Errorf("登陆日志保存失败，错误：%v", err.Error())
	}
}
