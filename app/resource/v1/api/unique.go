package api

import (
	"openops/app/resource/models"
	"openops/pkg/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/response"
)

// ModelUniqueByIdList 通过源模型ID获取唯一约束列表
func ModelUniqueByIdList(c *gin.Context) {
	var (
		err     error
		list    []*models.ModelUnique
		modelId = c.Param("id")
	)

	err = db.Orm().Model(&models.ModelUnique{}).Where("model_id =?", modelId).Find(&list).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelUniqueError)
		return
	}

	response.OK(c, nil, "")
}

// CreateModelUnique 创建唯一约束
func CreateModelUnique(c *gin.Context) {
	var (
		err    error
		unique *models.ModelUnique
		count  int64
	)

	err = c.ShouldBindJSON(&unique)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// model_id、field_id、type、title 唯一
	err = db.Orm().Model(&models.ModelUnique{}).
		Where("model_id = ? AND field_id = ? AND type = ? AND title = ?", unique.ModelId, unique.FieldId, unique.Type, unique.Title).
		Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelUniqueError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.ModelUniqueExistError)
		return
	}

	err = db.Orm().Create(&unique).Error
	if err != nil {
		response.Error(c, err, respstatus.CreateModelUniqueError)
		return
	}

	response.OK(c, unique, "")
}

// UpdateModelUnique 更新唯一约束
func UpdateModelUnique(c *gin.Context) {
	var (
		err      error
		unique   *models.ModelUnique
		uniqueId = c.Param("id")
		count    int64
	)

	err = c.ShouldBindJSON(&unique)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// model_id、field_id、type、title 唯一，排除自身
	err = db.Orm().Model(&models.ModelUnique{}).
		Where("model_id =? AND field_id =? AND type =? AND title =? AND id !=?", unique.ModelId, unique.FieldId, unique.Type, unique.Title, uniqueId).
		Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelUniqueError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.ModelUniqueExistError)
		return
	}

	err = db.Orm().Model(&models.ModelUnique{}).Where("id =?", uniqueId).Updates(map[string]interface{}{
		"title":    unique.Title,
		"type":     unique.Type,
		"field_id": unique.FieldId,
		"model_id": unique.ModelId,
		"desc":     unique.Desc,
	}).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateModelUniqueError)
		return
	}

	response.OK(c, unique, "")
}

// DeleteModelUnique 删除唯一约束
func DeleteModelUnique(c *gin.Context) {
	var (
		err      error
		uniqueId = c.Param("id")
	)

	err = db.Orm().Delete(&models.ModelUnique{}, "id =?", uniqueId).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteModelUniqueError)
		return
	}

	response.OK(c, nil, "")
}
