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

// LogicHandleListById 获取指定资源的逻辑处理列表
func LogicHandleListById(c *gin.Context) {
	var (
		err          error
		logicHandles []models.LogicHandle
		query        struct {
			LogicResource int    `form:"logic_resource"`
			Name          string `form:"name"`
			Remarks       string `form:"remarks"`
		}
	)

	err = c.ShouldBindQuery(&query)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	dbConn := db.Orm().Model(&models.LogicHandle{})

	if query.Name != "" {
		dbConn = dbConn.Where("name like ?", "%"+query.Name+"%")
	}

	if query.Remarks != "" {
		dbConn = dbConn.Where("remarks like ?", "%"+query.Remarks+"%")
	}

	if query.LogicResource != 0 {
		dbConn = dbConn.Where("logic_resource = ?", query.LogicResource)
	}

	err = dbConn.Find(&logicHandles).Error
	if err != nil {
		response.Error(c, err, respstatus.GetLogicHandleError)
		return
	}

	response.OK(c, logicHandles, "")
}

func CreateLogicHandle(c *gin.Context) {
	var (
		err         error
		logicHandle models.LogicHandle
		count       int64
	)

	err = c.ShouldBindJSON(&logicHandle)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// name 和 logic_resource 联合唯一
	err = db.Orm().Model(&models.LogicHandle{}).Where("name = ? AND logic_resource = ?", logicHandle.Name, logicHandle.LogicResource).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetLogicHandleError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.LogicHandleExistError)
		return
	}

	err = db.Orm().Create(&logicHandle).Error
	if err != nil {
		response.Error(c, err, respstatus.CreateLogicHandleError)
		return
	}

	response.OK(c, nil, "")
}

func UpdateLogicHandle(c *gin.Context) {
	var (
		err         error
		logicHandle models.LogicHandle
		count       int64
		id          = c.Param("id")
	)

	err = c.ShouldBindJSON(&logicHandle)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	err = db.Orm().Model(&models.LogicHandle{}).Where("name = ? AND logic_resource = ? and id != ?", logicHandle.Name, logicHandle.LogicResource, id).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetLogicHandleError)
		return
	}

	if count > 0 {
		response.Error(c, err, respstatus.LogicHandleExistError)
		return
	}

	err = db.Orm().Model(&models.LogicHandle{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":    logicHandle.Name,
		"remarks": logicHandle.Remarks,
	}).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateLogicHandleError)
		return
	}

	response.OK(c, nil, "")
}

func DeleteLogicHandle(c *gin.Context) {
	var (
		err error
		id  = c.Param("id")
	)

	err = db.Orm().Where("id = ?", id).Delete(&models.LogicHandle{}).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteLogicHandleError)
		return
	}

	response.OK(c, nil, "")
}
