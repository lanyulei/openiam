package apis

import (
	"errors"
	"fmt"
	"openiam/app/system/models"
	"openiam/common/middleware/permission"
	commonModels "openiam/common/models"
	"openiam/pkg/tools/common"
	"openiam/pkg/tools/respstatus"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	dbConn "github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
)

// ApiList 接口列表
func ApiList(c *gin.Context) {
	var (
		err     error
		apiList []*struct {
			models.Api
			GroupName string `json:"group_name"`
		}
		result interface{}
	)

	db := dbConn.Orm().Model(&models.Api{})

	fields := c.DefaultQuery("fields", "")
	value := c.DefaultQuery("value", "")
	if fields != "" && value != "" {
		db = db.Where(fmt.Sprintf("%s like ?", fields), "%"+value+"%")
	}

	group := c.DefaultQuery("group", "")
	if group != "" {
		groupKey := `"group"`
		if viper.GetString("db.type") == string(commonModels.DBTypeMySQL) {
			groupKey = "`group`"
		}
		db = db.Where(`? = ?`, groupKey, group)
		if group != "0" {
			db = db.Select("system_api_group.name as group_name, system_api.*").
				Joins(common.AddQuotesToSQLTableNames("left join system_api_group on system_api_group.id = system_api.group"))
		}
	}

	noForensics := c.DefaultQuery("no_forensics", "all")
	if noForensics != "all" {
		verifyPerm := false
		if noForensics == "true" {
			verifyPerm = true
		}
		db = db.Where(`"no_forensics" = ?`, verifyPerm)
	}

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: db,
	}, &apiList)
	if err != nil {
		response.Error(c, err, respstatus.ApiListError)
		return
	}
	response.OK(c, result, "")
}

// SaveApi 保存接口
func SaveApi(c *gin.Context) {
	var (
		err             error
		api, currentApi models.Api
	)

	err = c.ShouldBind(&api)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	db := dbConn.Orm().Model(&models.Api{})

	if api.Id != 0 {
		db = db.Where("id = ?", api.Id)

		err = dbConn.Orm().Where("id = ?", api.Id).Find(&currentApi).Error
		if err != nil {
			response.Error(c, err, respstatus.GetApiError)
			return
		}
		api.CreatedAt = currentApi.CreatedAt
	}

	err = db.Save(&api).Error
	if err != nil {
		response.Error(c, err, respstatus.SaveApiError)
		return
	}

	response.OK(c, "", "")
}

// DeleteApi 删除接口
func DeleteApi(c *gin.Context) {
	var (
		err          error
		apiId        string
		emnuApiCount int64
	)

	apiId = c.Param("id")

	// 查询是否有菜单绑定了接口
	err = dbConn.Orm().Model(&models.MenuApi{}).Where("api = ?", apiId).Count(&emnuApiCount).Error
	if err != nil {
		response.Error(c, nil, respstatus.GetApiMenuError)
		return
	}

	if emnuApiCount > 0 {
		response.Error(c, nil, respstatus.ApiUsedError)
		return
	}

	err = dbConn.Orm().Delete(&models.Api{}, apiId).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteApiError)
		return
	}

	response.OK(c, "", "")
}

func UpdateApiNoForensics(c *gin.Context) {
	var (
		err    error
		params struct {
			NoForensics bool  `json:"no_forensics"`
			Ids         []int `json:"ids"`
		}
	)

	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	if len(params.Ids) == 0 {
		response.Error(c, errors.New("ids 不能为空"), respstatus.InvalidParameterError)
		return
	}

	err = dbConn.Orm().Model(&models.Api{}).Where("id in (?)", params.Ids).Update("no_forensics", params.NoForensics).Error
	if err != nil {
		response.Error(c, err, respstatus.SaveApiError)
		return
	}

	// 获取不需要的验证的接口
	err = permission.SetNoForensicsAPIList()
	if err != nil {
		response.Error(c, err, respstatus.SaveApiError)
		return
	}

	response.OK(c, "", "")
}

func BatchCreateApi(c *gin.Context) {
	var (
		err                          error
		api, createApis, currentApis []*models.Api
		apiMap                       = make(map[string]struct{})
	)

	err = c.ShouldBindJSON(&api)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	// 获取当前所有接口
	err = dbConn.Orm().Model(&models.Api{}).Find(&currentApis).Error
	if err != nil {
		response.Error(c, err, respstatus.GetApiError)
		return
	}

	for _, v := range currentApis {
		apiMap[v.URL+v.Method] = struct{}{}
	}

	for _, v := range api {
		if _, ok := apiMap[v.URL+v.Method]; ok {
			continue
		}
		createApis = append(createApis, v)
	}

	err = dbConn.Orm().CreateInBatches(&createApis, 100).Error
	if err != nil {
		response.Error(c, err, respstatus.SaveApiError)
		return
	}

	response.OK(c, "", "")
}
