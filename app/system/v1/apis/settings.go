package apis

import (
	"github.com/spf13/viper"
	"openiam/app/system/models"
	commonModels "openiam/common/models"
	"openiam/pkg/tools/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/response"
)

/*
  @Author : lanyulei
  @Desc :
*/

// GetSettings
// @Description: 获取所有配置
// @param c
func GetSettings(c *gin.Context) {

	var (
		err      error
		settings []models.Settings
		result   = make(map[string]models.Settings)
		query    struct {
			Key string `form:"key"`
		}
	)

	err = c.ShouldBindQuery(&query)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	dbConn := db.Orm().Model(&models.Settings{})

	if query.Key != "" {
		keyValue := "\"key\""
		if viper.GetString("db.type") == string(commonModels.DBTypeMySQL) {
			keyValue = "`key`"
		}
		dbConn = dbConn.Where("? = ?", keyValue, query.Key)
	}

	err = dbConn.Find(&settings).Error
	if err != nil {
		response.Error(c, err, respstatus.GetSettingsError)
		return
	}

	for _, v := range settings {
		result[v.Key] = v
	}

	response.OK(c, result, "")
}

// UpdateSettings
// @Description: 更新配置
// @param c
func UpdateSettings(c *gin.Context) {
	var (
		err      error
		settings []models.Settings
		count    int64
	)

	err = c.ShouldBindJSON(&settings)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParameterError)
		return
	}

	for _, s := range settings {
		keyValue := "\"key\""
		if viper.GetString("db.type") == string(commonModels.DBTypeMySQL) {
			keyValue = "`key`"
		}

		err = db.Orm().Model(&models.Settings{}).Where("? = ?", keyValue, s.Key).Count(&count).Error
		if err != nil {
			response.Error(c, err, respstatus.GetSettingsError)
			return
		}

		if count > 0 {
			err = db.Orm().Model(&models.Settings{}).Where("? = ?", keyValue, s.Key).Update("content", s.Content).Error
			if err != nil {
				response.Error(c, err, respstatus.UpdateSettingsError)
				return
			}
		} else {
			err = db.Orm().Create(&s).Error
			if err != nil {
				response.Error(c, err, respstatus.UpdateSettingsError)
				return
			}
		}
	}

	response.OK(c, settings, "")
}
