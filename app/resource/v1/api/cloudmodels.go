package api

import (
	"openops/app/resource/models"
	"openops/pkg/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/response"
)

/*
  @Author : lanyulei
  @Desc :
*/

func GetCloudModels(c *gin.Context) {
	var (
		err         error
		cloudModels []struct {
			LogicResourceLabel string `json:"logic_resource_label"`
			LogicResourceValue string `json:"logic_resource_value"`
			models.CloudModels
		}
		query struct {
			CloudAccountId int `form:"cloud_account_id"`
		}
	)

	err = c.ShouldBindQuery(&query)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	dbConn := db.Orm().Model(&models.CloudModels{}).
		Select("cmdb_logic_resource.label as logic_resource_label, cmdb_logic_resource.value as logic_resource_value, cmdb_cloud_models.*").
		Joins("left join cmdb_logic_resource on cmdb_logic_resource.id = cmdb_cloud_models.logic_resource")

	if query.CloudAccountId != 0 {
		dbConn = dbConn.Where("cmdb_cloud_models.cloud_account_id = ?", query.CloudAccountId)
	}

	err = dbConn.Find(&cloudModels).Error
	if err != nil {
		response.Error(c, err, respstatus.GetCloudModelsError)
		return
	}

	response.OK(c, cloudModels, "")
}

func CreateCloudModel(c *gin.Context) {
	var (
		err        error
		cloudModel models.CloudModels
		count      int64
	)

	err = c.ShouldBindJSON(&cloudModel)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	err = db.Orm().Model(&models.CloudModels{}).
		Where("cloud_account_id = ? AND model_id = ? AND logic_resource = ?",
			cloudModel.CloudAccountId,
			cloudModel.ModelId,
			cloudModel.LogicResource,
		).
		Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetCloudModelsError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.CloudModelExistError)
		return
	}

	err = db.Orm().Create(&cloudModel).Error
	if err != nil {
		response.Error(c, err, respstatus.CreateCloudModelError)
		return
	}

	response.OK(c, "", "")
}

func UpdateCloudModel(c *gin.Context) {
	var (
		err        error
		cloudModel models.CloudModels
		modelId    = c.Param("id")
		count      int64
	)

	err = c.ShouldBindJSON(&cloudModel)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	err = db.Orm().Model(&models.CloudModels{}).
		Where("cloud_account_id = ? AND model_id = ? AND logic_resource = ? AND id != ?",
			cloudModel.CloudAccountId,
			cloudModel.ModelId,
			cloudModel.LogicResource,
			modelId,
		).
		Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetCloudModelsError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.CloudModelExistError)
		return
	}

	err = db.Orm().Model(&models.CloudModels{}).Where("id = ?", modelId).Updates(map[string]interface{}{
		"model_id":       cloudModel.ModelId,
		"logic_resource": cloudModel.LogicResource,
		"logic_handle":   cloudModel.LogicHandle,
	}).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateCloudModelError)
		return
	}

	response.OK(c, "", "")
}

func DeleteCloudModel(c *gin.Context) {
	modelId := c.Param("id")

	err := db.Orm().Delete(&models.CloudModels{}, "id = ?", modelId).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteCloudModelError)
		return
	}

	response.OK(c, "", "")
}
