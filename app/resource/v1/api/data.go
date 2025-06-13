package api

import (
	"openops/app/resource/models"
	"openops/pkg/respstatus"
	"openops/pkg/server"

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

	dbConn := db.Orm().Model(&models.Data{}).Where("model_id = ?", modelId)

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

	err = server.VerifyData(models.VerifyDataStatusCreate, &data)
	if err != nil {
		response.Error(c, err, respstatus.VerifyDataError)
		return
	}

	err = db.Orm().Create(&data).Error
	if err != nil {
		response.Error(c, err, respstatus.CreateDataError)
		return
	}

	response.OK(c, data, "")
}

// UpdateData 更新
func UpdateData(c *gin.Context) {
	var (
		err  error
		data models.Data
		id   = c.Param("id")
	)

	if err = c.ShouldBindJSON(&data); err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	data.Id = id

	err = server.VerifyData(models.VerifyDataStatusUpdate, &data)
	if err != nil {
		response.Error(c, err, respstatus.VerifyDataError)
		return
	}

	err = db.Orm().Model(&models.Data{}).Where("id = ?", id).Updates(map[string]interface{}{
		"model_id": data.ModelId,
		"data":     data.Data,
		"status":   data.Status,
	}).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateDataError)
		return
	}

	response.OK(c, data, "")
}

// BatchDeleteData 批量删除数据
func BatchDeleteData(c *gin.Context) {
	var (
		err  error
		data []string
	)

	err = c.ShouldBindJSON(&data)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	if len(data) == 0 {
		response.Error(c, nil, respstatus.InvalidParamsError)
		return
	}

	err = db.Orm().Model(&models.Data{}).Where("id IN ?", data).Delete(&models.Data{}).Error
	if err != nil {
		response.Error(c, err, respstatus.BatchDeleteDataError)
		return
	}

	response.OK(c, nil, "")
}

// DataDetails 获取数据详情
func DataDetails(c *gin.Context) {
	var (
		err  error
		id   = c.Param("id")
		data struct {
			models.Data
			ModelName string `json:"model_name" gorm:"-"`
		}
		modelInfo models.Model
	)

	err = db.Orm().Model(&models.Data{}).Where("id = ?", id).First(&data).Error
	if err != nil {
		response.Error(c, err, respstatus.GetDataDetailsError)
		return
	}

	err = db.Orm().Model(&models.Model{}).Where("id = ?", data.ModelId).First(&modelInfo).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelError)
		return
	}

	data.ModelName = modelInfo.Name

	response.OK(c, data, "")
}
