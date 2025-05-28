package api

import (
	"openops/app/resource/models"
	"openops/pkg/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
)

// FieldGroupList 分页获取字段分组列表
func FieldGroupList(c *gin.Context) {
	var (
		err    error
		list   []*models.FieldGroup
		result interface{}
	)

	dbConn := db.Orm().Model(&models.FieldGroup{})

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: dbConn,
	}, &list)
	if err != nil {
		response.Error(c, err, respstatus.FieldGroupListError)
		return
	}

	response.OK(c, result, "")
}
