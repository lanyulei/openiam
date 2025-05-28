package api

import (
	"openops/app/resource/models"
	"openops/pkg/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/response"
)

// ModelRelationBySourceModelIdList 获取模型关系列表
func ModelRelationBySourceModelIdList(c *gin.Context) {
	var (
		err       error
		relations []models.ModelRelation
		sourceId  = c.Param("sourceModelId")
	)

	if err = db.Orm().Where("source_model_id =?", sourceId).Find(&relations).Error; err != nil {
		response.Error(c, err, respstatus.GetModelRelationError)
		return
	}

	response.OK(c, relations, "")
}

// CreateModelRelation 创建模型关系
func CreateModelRelation(c *gin.Context) {
	var (
		err      error
		relation models.ModelRelation
		count    int64
	)

	if err = c.ShouldBindJSON(&relation); err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// source_model_id、target_model_id、type、constraint 联合唯一
	if err = db.Orm().
		Where("source_model_id =? AND target_model_id =? AND type =? AND constraint =?",
			relation.SourceModelId, relation.TargetModelId, relation.Type, relation.Constraint).
		Count(&count).Error; err != nil {
		response.Error(c, err, respstatus.GetModelRelationError)
		return
	}

	if err = db.Orm().Create(&relation).Error; err != nil {
		response.Error(c, err, respstatus.CreateModelRelationError)
		return
	}

	response.OK(c, relation, "")
}

// UpdateModelRelation 更新模型关系
func UpdateModelRelation(c *gin.Context) {
	var (
		err        error
		relation   models.ModelRelation
		count      int64
		relationId = c.Param("id")
	)

	if err = c.ShouldBindJSON(&relation); err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// source_model_id、target_model_id、type、constraint 联合唯一, 排除当前 id
	if err = db.Orm().
		Where("source_model_id =? AND target_model_id =? AND type =? AND constraint =? AND id !=?",
			relation.SourceModelId, relation.TargetModelId, relation.Type, relation.Constraint, relationId).
		Count(&count).Error; err != nil {
		response.Error(c, err, respstatus.GetModelRelationError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.ModelRelationExistError)
		return
	}

	err = db.Orm().Model(&models.ModelRelation{}).Updates(map[string]interface{}{
		"source_model_id": relation.SourceModelId,
		"target_model_id": relation.TargetModelId,
		"type":            relation.Type,
		"constraint":      relation.Constraint,
		"desc":            relation.Desc,
	}).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateModelRelationError)
		return
	}

	response.OK(c, relation, "")
}

// DeleteModelRelation 删除模型关系
func DeleteModelRelation(c *gin.Context) {
	var (
		err        error
		relationId = c.Param("id")
	)

	if err = db.Orm().Where("id = ?", relationId).Delete(&models.ModelRelation{}).Error; err != nil {
		response.Error(c, err, respstatus.DeleteModelRelationError)
		return
	}

	response.OK(c, nil, "")
}
