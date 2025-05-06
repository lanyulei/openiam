package apis

import (
	"openiam/app/system/models"
	"openiam/pkg/tools/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
)

/*
  @Author : lanyulei
  @Desc :
*/

func AppGroupList(c *gin.Context) {
	var (
		err    error
		list   []*models.AppGroup
		result interface{}
		name   string
	)

	dbConn := db.Orm().Model(&models.AppGroup{})

	name = c.DefaultQuery("name", "")
	if name != "" {
		dbConn = dbConn.Where("name like ?", "%"+name+"%")
	}

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: dbConn,
	}, &list)
	if err != nil {
		response.Error(c, err, respstatus.GetAppGroupError)
		return
	}

	response.OK(c, result, "")
}

func CreateAppGroup(c *gin.Context) {
	var (
		err      error
		appGroup models.AppGroup
		count    int64
	)

	if err = c.ShouldBindJSON(&appGroup); err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	// 名称存在则不创建
	if err = db.Orm().Model(&models.AppGroup{}).Where("name = ?", appGroup.Name).Count(&count).Error; err != nil {
		response.Error(c, err, respstatus.GetAppGroupError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.AppGroupExistError)
		return
	}

	if err = db.Orm().Create(&appGroup).Error; err != nil {
		response.Error(c, err, respstatus.CreateAppGroupError)
		return
	}

	response.OK(c, appGroup, "")
}

func UpdateAppGroup(c *gin.Context) {
	var (
		err             error
		appGroup        models.AppGroup
		count           int64
		appGroupId      string
		currentAppGroup models.AppGroup
		updateMap       = make(map[string]interface{})
	)

	appGroupId = c.Param("id")

	if err = c.ShouldBindJSON(&appGroup); err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	// 判断名称是否存在变化
	if err = db.Orm().Model(&models.AppGroup{}).Where("id = ?", appGroupId).Find(&currentAppGroup).Error; err != nil {
		response.Error(c, err, respstatus.GetAppGroupError)
		return
	}

	updateMap["remarks"] = appGroup.Remarks
	updateMap["sort"] = appGroup.Sort
	if appGroup.Name != currentAppGroup.Name {
		updateMap["name"] = appGroup.Name
		// 名称存在则不更新
		if err = db.Orm().Model(&models.AppGroup{}).Where("name = ?", appGroup.Name).Count(&count).Error; err != nil {
			response.Error(c, err, respstatus.GetAppGroupError)
			return
		}
		if count > 0 {
			response.Error(c, err, respstatus.AppGroupExistError)
			return
		}
	}

	// 更新应用组
	if err = db.Orm().Model(&models.AppGroup{}).Where("id = ?", appGroupId).Updates(updateMap).Error; err != nil {
		response.Error(c, err, respstatus.UpdateAppGroupError)
		return
	}

	response.OK(c, appGroup, "")
}

func DeleteAppGroup(c *gin.Context) {
	var (
		err        error
		appCount   int64
		appGroupId string
	)

	appGroupId = c.Param("id")

	// 应用组下存在应用则不允许删除
	if err = db.Orm().Model(&models.App{}).Where("app_group_id = ?", appGroupId).Count(&appCount).Error; err != nil {
		response.Error(c, err, respstatus.GetAppError)
		return
	}

	if appCount > 0 {
		response.Error(c, err, respstatus.AppGroupHasAppError)
		return
	}

	// 删除应用组
	if err = db.Orm().Delete(&models.AppGroup{}, "id = ?", appGroupId).Error; err != nil {
		response.Error(c, err, respstatus.DeleteAppGroupError)
		return
	}

	response.OK(c, "", "")
}
