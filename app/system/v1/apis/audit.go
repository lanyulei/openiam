package apis

import (
	"openiam/app/system/models"
	commonModels "openiam/common/models"
	"openiam/pkg/tools/respstatus"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
)

/*
  @Author : lanyulei
  @Desc :
*/

func AuditList(c *gin.Context) {
	var (
		err    error
		list   []*models.Audit
		result interface{}
	)

	SearchParams := map[string]map[string]interface{}{
		"like": pagination.RequestParams(c),
	}

	dbConn := db.Orm().Model(&models.Audit{})
	if viper.GetString("db.type") == string(commonModels.DBTypeMySQL) {
		dbConn.Select("id, username, path, method, browser, ip, `system`, create_time")
	} else if viper.GetString("db.type") == string(commonModels.DBTypePostgres) {
		dbConn = dbConn.Select("id, username, path, method, browser, ip, \"system\", create_time")
	}

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: dbConn,
	}, &list, SearchParams)
	if err != nil {
		response.Error(c, err, respstatus.GetAuditLogError)
		return
	}

	response.OK(c, result, "")
}

func DeleteAudit(c *gin.Context) {
	var (
		err error
		id  string
	)

	id = c.Param("id")

	err = db.Orm().Where("id = ?", id).Delete(&models.Audit{}).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteAuditLogError)
		return
	}

	response.OK(c, id, "")
}

func GetAuditInfo(c *gin.Context) {
	var (
		err  error
		id   string
		info models.Audit
	)

	id = c.Param("id")

	err = db.Orm().Where("id = ?", id).First(&info).Error
	if err != nil {
		response.Error(c, err, respstatus.GetAuditLogError)
		return
	}

	response.OK(c, info, "")
}
