package api

import (
	"openops/app/resource/models"
	"openops/pkg/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
)

// ModelGroupList 分页展示模型分组列表
func ModelGroupList(c *gin.Context) {
	var (
		err    error
		list   []*models.ModelGroup
		result interface{}
	)

	dbConn := db.Orm().Model(&models.ModelGroup{})

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: dbConn,
	}, &list)
	if err != nil {
		response.Error(c, err, respstatus.UserListError)
		return
	}

	response.OK(c, result, "")
}

// CreateModelGroup 创建模型分组
func CreateModelGroup(c *gin.Context) {
	var (
		err   error
		group models.ModelGroup
		count int64
	)

	err = c.ShouldBindJSON(&group)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// 分组名称唯一性校验
	err = db.Orm().Model(&models.ModelGroup{}).Where("name = ?", group.Name).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelGroupError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.ModelGroupNameExistError)
		return
	}

	err = db.Orm().Create(&group).Error
	if err != nil {
		response.Error(c, err, respstatus.CreateModelGroupError)
		return
	}

	response.OK(c, group, "")
}

// UpdateModelGroup 更新模型分组
func UpdateModelGroup(c *gin.Context) {
	var (
		err          error
		group        models.ModelGroup
		modelGroupId = c.Param("id")
		count        int64
	)

	err = c.ShouldBindJSON(&group)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// 分组名称唯一性校验，排除当前分组
	err = db.Orm().Model(&models.ModelGroup{}).Where("name = ? AND id != ?", group.Name, modelGroupId).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelGroupError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.ModelGroupNameExistError)
		return
	}

	err = db.Orm().Model(&models.ModelGroup{}).Where("id = ?", modelGroupId).Updates(map[string]interface{}{
		"name":  group.Name,
		"desc":  group.Desc,
		"order": group.Order,
	}).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateModelGroupError)
		return
	}

	response.OK(c, group, "")
}

// DeleteModelGroup 删除模型分组
func DeleteModelGroup(c *gin.Context) {
	var (
		err          error
		modelGroupId = c.Param("id")
		count        int64
	)

	// 检查是否有模型绑定该分组
	err = db.Orm().Model(&models.Model{}).Where("group_id =?", modelGroupId).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetModelGroupError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.ModelGroupHasModelError)
		return
	}

	err = db.Orm().Model(&models.ModelGroup{}).Where("id = ?", modelGroupId).Delete(&models.ModelGroup{}).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteModelGroupError)
		return
	}

	response.OK(c, nil, "")
}
