package api

import (
	"openops/app/resource/models"
	"openops/pkg/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
)

// GetModels 获取模型列表
func GetModels(c *gin.Context) {
	var (
		err  error
		list []*struct {
			models.ModelGroup
			Models []*models.Model `json:"models" gorm:"-"`
		}
		modelList []*models.Model
		modelMap  = make(map[string][]*models.Model)
	)

	err = db.Orm().Model(&models.ModelGroup{}).
		Order(`"order" asc`).
		Find(&list).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelGroupError)
		return
	}

	err = db.Orm().Model(&models.Model{}).
		Order(`"order" asc`).
		Find(&modelList).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelError)
		return
	}

	for _, model := range modelList {
		modelMap[model.GroupId] = append(modelMap[model.GroupId], model)
	}

	for _, group := range list {
		group.Models = modelMap[group.Id]
	}

	response.OK(c, list, "")
}

// ModelList 分页展示模型列表
func ModelList(c *gin.Context) {
	var (
		err    error
		list   []*models.Model
		result interface{}
	)

	dbConn := db.Orm().Model(&models.Model{})

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: dbConn,
	}, &list)
	if err != nil {
		response.Error(c, err, respstatus.ModelListError)
		return
	}

	response.OK(c, result, "")
}

// CreateModel 创建模型
func CreateModel(c *gin.Context) {
	var (
		err   error
		model models.Model
		count int64
	)

	err = c.ShouldBindJSON(&model)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// 模型名称唯一性校验
	err = db.Orm().Model(&models.Model{}).Where("name = ?", model.Name).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelError)
		return
	}
	if count > 0 {
		response.Error(c, err, respstatus.ModelNameExistError)
		return
	}

	err = db.Orm().Create(&model).Error
	if err != nil {
		response.Error(c, err, respstatus.CreateModelError)
		return
	}

	response.OK(c, model, "")
}

// UpdateModel 更新模型
func UpdateModel(c *gin.Context) {
	var (
		err     error
		model   models.Model
		modelId = c.Param("id")
		count   int64
	)

	err = c.ShouldBindJSON(&model)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// 模型名称唯一性校验，排除当前模型
	err = db.Orm().Model(&models.Model{}).Where("name = ? AND id != ?", model.Name, modelId).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelError)
		return
	}
	if count > 0 {
		response.Error(c, err, respstatus.ModelNameExistError)
		return
	}

	err = db.Orm().Model(&models.Model{}).Where("id = ?", modelId).Updates(map[string]interface{}{
		"name":     model.Name,
		"icon":     model.Icon,
		"status":   model.Status,
		"desc":     model.Desc,
		"group_id": model.GroupId,
		"order":    model.Order,
	}).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateModelError)
		return
	}

	response.OK(c, model, "")
}

// DeleteModel 删除模型
func DeleteModel(c *gin.Context) {
	var (
		err                               error
		modelId                           = c.Param("id")
		count, fieldGroupCount, dataCount int64
	)

	// 检查模型是否绑定了字段分组
	err = db.Orm().Model(&models.FieldGroup{}).Where("model_id = ?", modelId).Count(&fieldGroupCount).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelFieldGroupError)
		return
	}

	if fieldGroupCount > 0 {
		// 如果模型绑定了字段分组，不允许删除
		response.Error(c, err, respstatus.ModelHasFieldGroupError)
		return
	}

	// 检查模型是否绑定了字段
	err = db.Orm().Model(&models.Field{}).Where("model_id =?", modelId).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelFieldError)
		return
	}

	if count > 0 {
		// 如果模型绑定了字段，不允许删除
		response.Error(c, err, respstatus.ModelHasFieldError)
		return
	}

	// 检查模型是否绑定了数据
	err = db.Orm().Model(&models.Data{}).Where("model_id = ?", modelId).Count(&dataCount).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelDataError)
		return
	}

	if dataCount > 0 {
		// 如果模型绑定了数据，不允许删除
		response.Error(c, err, respstatus.ModelHasDataError)
		return
	}

	err = db.Orm().Where("id = ?", modelId).Delete(&models.Model{}).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteModelError)
		return
	}

	response.OK(c, nil, "")
}
