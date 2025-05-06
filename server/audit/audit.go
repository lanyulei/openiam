package audit

import (
	"encoding/json"
	"openiam/app/system/models"
	"openiam/pkg/tools"

	"github.com/lanyulei/toolkit/logger"

	"github.com/lanyulei/toolkit/db"

	"github.com/gin-gonic/gin"

	"github.com/mssola/user_agent"
)

/*
  @Author : lanyulei
  @Desc :
*/

func Create(c *gin.Context, username string, data []byte) (err error) {
	var (
		query []byte
	)

	ua := user_agent.New(c.Request.UserAgent())
	browserName, browserVersion := ua.Browser()

	auditValue := models.Audit{
		Username: username,
		Path:     c.Request.URL.Path,
		Method:   c.Request.Method,
		Browser:  browserName + " " + browserVersion,
		IP:       tools.GetClientIP(c),
		System:   ua.OS(),
	}

	if string(data) != "" {
		auditValue.Data = data
	}

	// query
	query, err = json.Marshal(c.Request.URL.Query())
	if err != nil {
		logger.Errorf("failed to marshal query, error: %v", err.Error())
		return
	}
	if string(query) != "" {
		auditValue.Query = query
	}

	err = db.Orm().Create(&auditValue).Error
	if err != nil {
		logger.Errorf("add audit failure, err: %v", err.Error())
		return
	}

	return
}
