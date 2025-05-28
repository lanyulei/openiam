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

// CreateFieldGroup 创建字段分组
func CreateFieldGroup(c *gin.Context) {
	var (
		err   error
		req   models.FieldGroup
		count int64
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// name 必须唯一
	if err = db.Orm().Model(&models.FieldGroup{}).Where("name = ?", req.Name).Count(&count).Error; err != nil {
		response.Error(c, err, respstatus.GetFieldGroupError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.FieldGroupNameExistError)
		return
	}

	if err = db.Orm().Create(&req).Error; err != nil {
		response.Error(c, err, respstatus.CreateFieldGroupError)
		return
	}

	response.OK(c, req, "")
}

// UpdateFieldGroup 更新字段分组
func UpdateFieldGroup(c *gin.Context) {
	var (
		err          error
		req          models.FieldGroup
		count        int64
		fieldGroupId = c.Param("id")
	)

	if err = c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// name 必须唯一，排除自己
	if err = db.Orm().Model(&models.FieldGroup{}).Where("name = ? and id != ?", req.Name, fieldGroupId).Count(&count).Error; err != nil {
		response.Error(c, err, respstatus.GetFieldGroupError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.FieldGroupNameExistError)
		return
	}

	if err = db.Orm().Model(&models.FieldGroup{}).Where("id = ?", fieldGroupId).Updates(map[string]interface{}{
		"name":     req.Name,
		"desc":     req.Desc,
		"order":    req.Order,
		"model_id": req.ModelId,
	}).Error; err != nil {
		response.Error(c, err, respstatus.UpdateFieldGroupError)
		return
	}

	response.OK(c, req, "")
}

// DeleteFieldGroup 删除字段分组
func DeleteFieldGroup(c *gin.Context) {
	var (
		err          error
		fieldGroupId = c.Param("id")
		count        int64
	)

	// 字段绑定了字段分组，则不能删除
	if err = db.Orm().Model(&models.Field{}).Where("group_id =?", fieldGroupId).Count(&count).Error; err != nil {
		response.Error(c, err, respstatus.GetFieldGroupError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.FieldGroupHasBindError)
		return
	}

	if err = db.Orm().Delete(&models.FieldGroup{}, "id = ?", fieldGroupId).Error; err != nil {
		response.Error(c, err, respstatus.DeleteFieldGroupError)
		return
	}

	response.OK(c, nil, "")
}
