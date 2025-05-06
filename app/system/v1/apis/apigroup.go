package apis

import (
	"openiam/app/system/models"
	commonModels "openiam/common/models"
	"openiam/pkg/tools/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
	"github.com/spf13/viper"
)

// ApiGroupList 接口分组列表
func ApiGroupList(c *gin.Context) {
	var (
		err          error
		apiGroupList []*models.ApiGroup
		result       interface{}
	)

	SearchParams := map[string]map[string]interface{}{
		"like": pagination.RequestParams(c),
	}

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: db.Orm().Model(&models.ApiGroup{}).Order("sort"),
	}, &apiGroupList, SearchParams)
	if err != nil {
		response.Error(c, err, respstatus.ApiGroupListError)
		return
	}
	response.OK(c, result, "")
}

// SaveApiGroup 保存接口分组
func SaveApiGroup(c *gin.Context) {
	var (
		err                       error
		apiGroup, currentApiGroup models.ApiGroup
		apiGroupCount             int64
	)

	err = c.ShouldBind(&apiGroup)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	dbModels := db.Orm().Model(&models.ApiGroup{})

	if apiGroup.Id != 0 {
		dbModels = dbModels.Where("id = ?", apiGroup.Id)

		err = db.Orm().Where("id = ?", apiGroup.Id).Find(&currentApiGroup).Error
		if err != nil {
			response.Error(c, err, respstatus.GetApiGroupError)
			return
		}

		apiGroup.CreatedAt = currentApiGroup.CreatedAt
	} else {
		err = db.Orm().Model(&models.ApiGroup{}).
			Where(`"name" = ?`, apiGroup.Name).
			Count(&apiGroupCount).Error
		if err != nil {
			response.Error(c, err, respstatus.ApiGroupExistError)
			return
		}
	}

	err = dbModels.Save(&apiGroup).Error
	if err != nil {
		response.Error(c, err, respstatus.SaveApiGroupError)
		return
	}

	response.OK(c, "", "")
}

// DeleteApiGroup 删除接口分组
func DeleteApiGroup(c *gin.Context) {
	var (
		err        error
		apiCount   int64
		apiGroupId string
	)

	apiGroupId = c.Param("id")
	groupKey := `"group"`
	if viper.GetString("db.type") == string(commonModels.DBTypeMySQL) {
		groupKey = "`group`"
	}
	err = db.Orm().Model(&models.Api{}).Where(`? = ?`, groupKey, apiGroupId).Count(&apiCount).Error
	if err != nil {
		response.Error(c, err, respstatus.GetApiError)
		return
	}
	if apiCount > 0 {
		response.Error(c, err, respstatus.ApiGroupUsedError)
		return
	}

	err = db.Orm().Delete(&models.ApiGroup{}, apiGroupId).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteApiGroupError)
		return
	}

	response.OK(c, "", "")
}
