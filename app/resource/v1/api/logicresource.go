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

// LogicResourceList 获取逻辑资源列表
func LogicResourceList(c *gin.Context) {
	var (
		err           error
		logicResource []models.LogicResource
		query         struct {
			Label          string `form:"label"`
			CloudAccountId int    `form:"cloud_account_id"`
		}
		idList []int
	)

	err = c.ShouldBindQuery(&query)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	dbConn := db.Orm().Model(&models.LogicResource{})

	if query.Label != "" {
		dbConn = dbConn.Where("label like ?", "%"+query.Label+"%")
	}

	if query.CloudAccountId != 0 {
		err = db.Orm().Model(&models.CloudModels{}).Where("cloud_account_id = ?", query.CloudAccountId).Pluck("DISTINCT logic_resource", &idList).Error
		if err != nil {
			response.Error(c, err, respstatus.GetCloudModelError)
			return
		}
		if len(idList) > 0 {
			dbConn = dbConn.Where("id in (?)", idList)
		} else {
			dbConn = dbConn.Where("id in (0)")
		}
	}

	err = dbConn.Find(&logicResource).Error
	if err != nil {
		response.Error(c, err, respstatus.GetLogicResourceError)
		return
	}

	response.OK(c, logicResource, "")
}

func LogicResourceDetails(c *gin.Context) {
	var (
		err           error
		logicResource models.LogicResource
		id            = c.Param("id")
	)

	err = db.Orm().Where("id = ?", id).First(&logicResource).Error
	if err != nil {
		response.Error(c, err, respstatus.GetLogicResourceError)
		return
	}

	response.OK(c, logicResource, "")
}

// CreateLogicResource 创建逻辑资源
func CreateLogicResource(c *gin.Context) {
	var (
		err           error
		logicResource models.LogicResource
		count         int64
	)

	err = c.ShouldBindJSON(&logicResource)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// label 和 value 联合唯一
	err = db.Orm().Model(&models.LogicResource{}).Where("label = ? AND value = ?", logicResource.Label, logicResource.Value).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetLogicResourceError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.LogicResourceExistError)
		return
	}

	err = db.Orm().Create(&logicResource).Error
	if err != nil {
		response.Error(c, err, respstatus.CreateLogicResourceError)
		return
	}

	response.OK(c, nil, "")
}

// UpdateLogicResource 更新逻辑资源
func UpdateLogicResource(c *gin.Context) {
	var (
		err           error
		logicResource models.LogicResource
		count         int64
		id            = c.Param("id")
	)

	err = c.ShouldBindJSON(&logicResource)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	err = db.Orm().Model(&models.LogicResource{}).Where("label = ? AND value = ? and id != ?", logicResource.Label, logicResource.Value, id).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetLogicResourceError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.LogicResourceExistError)
		return
	}

	err = db.Orm().Model(&models.LogicResource{}).Where("id = ?", id).Updates(map[string]interface{}{
		"label": logicResource.Label,
		"value": logicResource.Value,
	}).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateLogicResourceError)
		return
	}

	response.OK(c, nil, "")
}

// DeleteLogicResource 删除逻辑资源
func DeleteLogicResource(c *gin.Context) {
	var (
		err   error
		id    = c.Param("id")
		count int64
	)

	// 未绑定逻辑操作，才可以删除
	err = db.Orm().Model(&models.LogicHandle{}).Where("logic_resource = ?", id).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetLogicHandleError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.LogicResourceBindError)
		return
	}

	err = db.Orm().Where("id = ?", id).Delete(&models.LogicResource{}).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteLogicResourceError)
		return
	}

	response.OK(c, nil, "")
}
