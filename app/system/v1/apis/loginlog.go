package apis

import (
	"openiam/app/system/models"
	"openiam/pkg/tools/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
)

// LoginLogList 登陆日志
func LoginLogList(c *gin.Context) {
	var (
		err    error
		list   []*models.LoginLog
		result interface{}
	)

	SearchParams := map[string]map[string]interface{}{
		"like": pagination.RequestParams(c),
	}

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: db.Orm().Model(&models.LoginLog{}),
	}, &list, SearchParams)
	if err != nil {
		response.Error(c, err, respstatus.LoginLogListError)
		return
	}
	response.OK(c, result, "")
}

// DeleteLoginLog 删除登陆
func DeleteLoginLog(c *gin.Context) {
	var (
		err        error
		loginLogId string
	)

	loginLogId = c.Param("id")

	err = db.Orm().Delete(&models.LoginLog{}, loginLogId).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteLoginLogError)
		return
	}

	response.OK(c, "", "")
}

// LoginLogInfo
// @Description: 获取当前登录用户的最近一次登录详情失败
// @param c
func LoginLogInfo(c *gin.Context) {
	var (
		err          error
		loginLogInfo models.LoginLog
	)

	err = db.Orm().
		Where("username = ?", c.GetString("username")).
		Order("id desc").
		Find(&loginLogInfo).Error
	if err != nil {
		response.Error(c, err, respstatus.LoginLogInfoError)
		return
	}

	response.OK(c, loginLogInfo, "")
}
