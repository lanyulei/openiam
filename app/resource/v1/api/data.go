package api

import (
	"openops/app/resource/models"
	"openops/pkg/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
)

func DataList(c *gin.Context) {
	var (
		err     error
		list    []*models.Data
		result  interface{}
		modelId = c.Param("id")
	)

	dbConn := db.Orm().Model(&models.Data{}).Where("model_id =?", modelId)

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: dbConn,
	}, &list)
	if err != nil {
		response.Error(c, err, respstatus.DataListError)
		return
	}

	response.OK(c, result, "")
}

// CreateData 创建
func CreateData(c *gin.Context) {
	var (
		err  error
		data models.Data
	)

	if err = c.ShouldBindJSON(&data); err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	response.OK(c, data, "")
}
