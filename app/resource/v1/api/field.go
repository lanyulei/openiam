package api

import (
	"openops/app/resource/models"
	"openops/pkg/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
)

// FieldList 分页查询
func FieldList(c *gin.Context) {
	var (
		err    error
		list   []*models.Field
		result interface{}
	)

	dbConn := db.Orm().Model(&models.Field{})

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: dbConn,
	}, &list)
	if err != nil {
		response.Error(c, err, respstatus.FieldListError)
		return
	}

	response.OK(c, result, "")
}

// CreateField 创建
func CreateField(c *gin.Context) {
	var (
		err      error
		field    models.Field
		count    int64
		keyCount int64
	)

	if err = c.ShouldBindJSON(&field); err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// model_id、name 联合唯一
	if err = db.Orm().Model(&models.Field{}).
		Where("model_id = ? AND name = ?", field.ModelId, field.Name).
		Count(&count).Error; err != nil {
		response.Error(c, err, respstatus.GetFieldError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.FieldExistError)
		return
	}

	// model_id、key 联合唯一
	if err = db.Orm().Model(&models.Field{}).
		Where("model_id = ? AND key = ?", field.ModelId, field.Key).
		Count(&keyCount).Error; err != nil {
		response.Error(c, err, respstatus.GetFieldError)
		return
	}
	if keyCount > 0 {
		response.Error(c, err, respstatus.FieldKeyExistError)
		return
	}

	if err = db.Orm().Create(&field).Error; err != nil {
		response.Error(c, err, respstatus.CreateFieldError)
		return
	}

	response.OK(c, field, "")
}

// UpdateField 更新
func UpdateField(c *gin.Context) {
	var (
		err      error
		field    models.Field
		count    int64
		keyCount int64
		fieldId  = c.Param("id")
	)

	if err = c.ShouldBindJSON(&field); err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// model_id、name 联合唯一，排除自己
	if err = db.Orm().Model(&models.Field{}).
		Where("model_id = ? AND name = ? AND id != ?", field.ModelId, field.Name, fieldId).
		Count(&count).Error; err != nil {
		response.Error(c, err, respstatus.GetFieldError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.FieldExistError)
		return
	}

	// model_id、key 联合唯一，排除自己
	if err = db.Orm().Model(&models.Field{}).
		Where("model_id =? AND key =? AND id!=?", field.ModelId, field.Key, fieldId).
		Count(&keyCount).Error; err != nil {
		response.Error(c, err, respstatus.GetFieldError)
		return
	}
	if keyCount > 0 {
		response.Error(c, err, respstatus.FieldKeyExistError)
		return
	}

	err = db.Orm().Model(&models.Field{}).
		Where("id = ?", fieldId).
		Updates(map[string]interface{}{
			"key":         field.Key,
			"name":        field.Name,
			"group_id":    field.GroupId,
			"type":        field.Type,
			"options":     field.Options,
			"is_edit":     field.IsEdit,
			"is_required": field.IsRequired,
			"is_list":     field.IsList,
			"placeholder": field.Placeholder,
			"desc":        field.Desc,
			"order":       field.Order,
			"model_id":    field.ModelId,
		}).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateFieldError)
		return
	}

	response.OK(c, field, "")
}

// DeleteField 删除
func DeleteField(c *gin.Context) {
	var (
		err     error
		fieldId = c.Param("id")
	)

	err = db.Orm().
		Where("id = ?", fieldId).
		Delete(&models.Field{}).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteFieldError)
		return
	}

	response.OK(c, nil, "")
}

// GetFieldsAndGroups 获取 field 和 field 的分组
func GetFieldsAndGroups(c *gin.Context) {
	var (
		err  error
		list []*struct {
			models.FieldGroup
			Fields []*models.Field `json:"fields" gorm:"-"`
		}
		fieldList []*models.Field
		fieldMap  = make(map[string][]*models.Field)
		modelId   = c.Param("id")
		query     struct {
			Name string `json:"name" form:"name"`
		}
	)

	if err = c.ShouldBindQuery(&query); err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// 获取 field 分组列表
	err = db.Orm().Model(&models.FieldGroup{}).
		Where("model_id = ?", modelId).
		Order(`"order" asc`).
		Find(&list).Error
	if err != nil {
		response.Error(c, err, respstatus.GetFieldGroupError)
		return
	}

	dbConn := db.Orm().Model(&models.Field{}).
		Where("model_id = ?", modelId).
		Order(`"order" asc`)

	if query.Name != "" {
		dbConn = dbConn.Where("name LIKE ?", "%"+query.Name+"%")
	}

	// 获取 field 列表
	err = dbConn.
		Find(&fieldList).Error
	if err != nil {
		response.Error(c, err, respstatus.GetFieldError)
		return
	}

	// 将 field 按分组 ID 分组
	for _, field := range fieldList {
		fieldMap[field.GroupId] = append(fieldMap[field.GroupId], field)
	}

	// 将分组后的 field 关联到对应的分组上
	for _, group := range list {
		group.Fields = fieldMap[group.Id]
	}

	response.OK(c, list, "")
}
