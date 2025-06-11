package api

import (
	"openops/app/resource/models"
	"openops/pkg/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
	"gorm.io/gorm"
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
		query     struct {
			Name string `json:"name" form:"name"`
		}
		modelIdList    []string
		modelDataCount []struct {
			ModelId string
			Count   int
		}
		modelDataCountMap = make(map[string]int)
	)

	if err = c.ShouldBindQuery(&query); err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	err = db.Orm().Model(&models.ModelGroup{}).
		Order(`"order" asc`).
		Find(&list).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelGroupError)
		return
	}

	dbConn := db.Orm().Model(&models.Model{}).
		Order(`"order" asc`)

	if query.Name != "" {
		dbConn = dbConn.Where("name LIKE ?", "%"+query.Name+"%")
	}

	err = dbConn.
		Find(&modelList).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelError)
		return
	}

	for _, model := range modelList {
		modelIdList = append(modelIdList, model.Id)
		modelMap[model.GroupId] = append(modelMap[model.GroupId], model)
	}

	for _, group := range list {
		group.Models = modelMap[group.Id]
	}

	// 获取所有模型的示例数据数量统计
	err = db.Orm().Model(&models.Data{}).
		Select("model_id, count(id) as count").
		Where("model_id in ?", modelIdList).
		Group("model_id").
		Find(&modelDataCount).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelError)
		return
	}

	// 模型数据数量统计
	for _, item := range modelDataCount {
		modelDataCountMap[item.ModelId] = item.Count
	}

	// 模型数据数量赋值
	for _, group := range list {
		for _, model := range group.Models {
			model.DataCount = modelDataCountMap[model.Id]
		}
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
		err       error
		modelId   = c.Param("id")
		dataCount int64
	)

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

	// gorm 的事务 Transaction
	err = db.Orm().Transaction(func(tx *gorm.DB) error {
		// 删除模型对应的字段分组
		err = tx.Where("model_id = ?", modelId).Delete(&models.FieldGroup{}).Error
		if err != nil {
			return err
		}

		// 删除模型对应的字段
		err = tx.Where("model_id =?", modelId).Delete(&models.Field{}).Error
		if err != nil {
			return err
		}

		err = tx.Where("id = ?", modelId).Delete(&models.Model{}).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		response.Error(c, err, respstatus.DeleteModelError)
		return
	}

	response.OK(c, nil, "")
}

// GetModel 获取模型详情
func GetModel(c *gin.Context) {
	var (
		err   error
		model struct {
			models.Model
			GroupName string `json:"group_name" gorm:"-"`
		}
		modelGroup models.ModelGroup
		modelId    = c.Param("id")
	)

	err = db.Orm().Where("id = ?", modelId).First(&model).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelError)
		return
	}

	err = db.Orm().Where("id =?", model.GroupId).First(&modelGroup).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelGroupError)
		return
	}

	model.GroupName = modelGroup.Name

	response.OK(c, model, "")
}
