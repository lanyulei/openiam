package api

import (
	"github.com/lanyulei/toolkit/pagination"
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
		err    error
		result interface{}
		list   []*models.LogicResource
	)

	dbConn := db.Orm().Model(&models.LogicResource{})

	name := c.Query("name")
	if name != "" {
		dbConn = dbConn.Or("name like ? or title like ?", "%"+name+"%", "%"+name+"%")
	}

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: dbConn,
	}, &list)
	if err != nil {
		response.Error(c, err, respstatus.LogicResourceListError)
		return
	}

	response.OK(c, result, "")
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

	// title 和 name 联合唯一
	err = db.Orm().Model(&models.LogicResource{}).Where("title = ? AND name = ?", logicResource.Title, logicResource.Name).Count(&count).Error
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

	err = db.Orm().Model(&models.LogicResource{}).Where("title = ? AND name = ? and id != ?", logicResource.Title, logicResource.Name, id).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetLogicResourceError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.LogicResourceExistError)
		return
	}

	err = db.Orm().Model(&models.LogicResource{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title": logicResource.Title,
		"name":  logicResource.Name,
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
