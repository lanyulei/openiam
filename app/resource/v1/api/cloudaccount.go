package api

import (
	"openops/app/resource/models"
	"openops/pkg/cloud/sync"
	"openops/pkg/cloud/types"
	"openops/pkg/crypto"
	"openops/pkg/respstatus"

	"github.com/gin-gonic/gin"
	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/pagination"
	"github.com/lanyulei/toolkit/response"
	"github.com/spf13/viper"
)

/*
  @Author : lanyulei
  @Desc :
*/

func CloudAccountList(c *gin.Context) {
	var (
		err              error
		cloudAccountList []*models.CloudAccount
		result           interface{}
		name             string
	)

	name = c.Query("name")

	dbConn := db.Orm().Model(&models.CloudAccount{})
	if name != "" {
		dbConn = dbConn.Where("name like ?", "%"+name+"%")
	}

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: dbConn,
	}, &cloudAccountList)
	if err != nil {
		response.Error(c, err, respstatus.GetCloudAccountError)
		return
	}

	for _, cloudAccount := range cloudAccountList {
		cloudAccount.ProviderName = types.CloudInfo[cloudAccount.Provider].Name
	}

	response.OK(c, result, "")
}

func CreateCloudAccount(c *gin.Context) {
	var (
		err          error
		cloudAccount struct {
			models.CloudAccount
			AccessKey string `json:"access_key" binging:"required"`
			SecretKey string `json:"secret_key" binging:"required"`
		}
		count int64
	)

	err = c.ShouldBindJSON(&cloudAccount)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// name 不存在则不能创建
	err = db.Orm().Model(&models.CloudAccount{}).Where("name = ?", cloudAccount.Name).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetCloudAccountError)
		return
	}
	if count > 0 {
		response.Error(c, err, respstatus.CloudAccountExistError)
		return
	}

	key := viper.GetString("aes.key")

	cloudAccount.AccessKey, err = crypto.AesEncryptCBC([]byte(key), []byte(cloudAccount.AccessKey))
	if err != nil {
		response.Error(c, err, respstatus.EncryptError)
		return
	}

	cloudAccount.SecretKey, err = crypto.AesEncryptCBC([]byte(key), []byte(cloudAccount.SecretKey))
	if err != nil {
		response.Error(c, err, respstatus.EncryptError)
		return
	}

	err = db.Orm().Create(&cloudAccount).Error
	if err != nil {
		response.Error(c, err, respstatus.CreateCloudAccountError)
		return
	}

	response.OK(c, nil, "")
}

func EditCloudAccount(c *gin.Context) {
	var (
		err                      error
		params, cloudAccountInfo models.CloudAccount
		count                    int64
		cloudAccountId           string
	)

	cloudAccountId = c.Param("id")

	err = db.Orm().Model(&models.CloudAccount{}).Where("id = ?", cloudAccountId).First(&cloudAccountInfo).Error
	if err != nil {
		response.Error(c, err, respstatus.GetCloudAccountError)
		return
	}

	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	// 更新数据的时候 name 不能重复
	err = db.Orm().Model(&models.CloudAccount{}).Where("name = ? and id != ?", params.Name, cloudAccountId).Count(&count).Error
	if err != nil {
		response.Error(c, err, respstatus.GetCloudAccountError)
		return
	}
	if count > 0 {
		response.Error(c, err, respstatus.CloudAccountExistError)
		return
	}

	key := viper.GetString("aes.key")

	if cloudAccountInfo.AccessKey != params.AccessKey {
		params.AccessKey, err = crypto.AesEncryptCBC([]byte(key), []byte(params.AccessKey))
		if err != nil {
			response.Error(c, err, respstatus.EncryptError)
			return
		}
	}

	if cloudAccountInfo.SecretKey != params.SecretKey {
		params.SecretKey, err = crypto.AesEncryptCBC([]byte(key), []byte(params.SecretKey))
		if err != nil {
			response.Error(c, err, respstatus.EncryptError)
			return
		}
	}

	err = db.Orm().Model(&models.CloudAccount{}).Where("id = ?", cloudAccountId).Updates(map[string]interface{}{
		"provider":    params.Provider,
		"name":        params.Name,
		"access_key":  params.AccessKey,
		"secret_key":  params.SecretKey,
		"available":   params.Available,
		"plugin_name": params.PluginName,
		"remarks":     params.Remarks,
	}).Error
	if err != nil {
		response.Error(c, err, respstatus.UpdateCloudAccountError)
		return
	}

	response.OK(c, nil, "")
}

func DeleteCloudAccount(c *gin.Context) {
	var (
		err                                error
		cloudAccountId                     string
		cloudModelsCount, cloudRegionCount int64
	)

	cloudAccountId = c.Param("id")

	// 存在绑定的资源则无法删除账号信息
	err = db.Orm().Model(&models.CloudModels{}).Where("cloud_account_id = ?", cloudAccountId).Count(&cloudModelsCount).Error
	if err != nil {
		response.Error(c, err, respstatus.GetCloudModelsError)
		return
	}

	if cloudModelsCount > 0 {
		response.Error(c, err, respstatus.CloudAccountBindResourceError)
		return
	}

	// 存在绑定的 Region 则无法删除账号信息
	err = db.Orm().Model(&models.CloudRegion{}).Where("cloud_account_id = ?", cloudAccountId).Count(&cloudRegionCount).Error
	if err != nil {
		response.Error(c, err, respstatus.GetCloudRegionError)
		return
	}

	if cloudRegionCount > 0 {
		response.Error(c, err, respstatus.CloudAccountBindRegionError)
		return
	}

	err = db.Orm().Model(&models.CloudAccount{}).Where("id = ?", cloudAccountId).Delete(&models.CloudAccount{}).Error
	if err != nil {
		response.Error(c, err, respstatus.DeleteCloudAccountError)
		return
	}

	response.OK(c, nil, "")
}

func CloudAccountCheckConnect(c *gin.Context) {
	// todo 修改为验证 插件的 连通性
	response.OK(c, nil, "")
}

func CloudAccountDetail(c *gin.Context) {
	var (
		err            error
		CloudAccountId string
		cloudAccount   models.CloudAccount
	)

	CloudAccountId = c.Param("id")

	err = db.Orm().Model(&models.CloudAccount{}).Where("id = ?", CloudAccountId).First(&cloudAccount).Error
	if err != nil {
		response.Error(c, err, respstatus.GetCloudAccountError)
		return
	}

	response.OK(c, cloudAccount, "")
}

func SyncCloudResource(c *gin.Context) {
	var (
		err    error
		params struct {
			CloudAccountId string `json:"cloud_account_id"`
		}
	)

	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.Error(c, err, respstatus.InvalidParamsError)
		return
	}

	err = db.Orm().Model(&models.CloudAccount{}).Where("id = ?", params.CloudAccountId).Updates(map[string]interface{}{
		"sync_status":  models.SyncStatusRunning,
		"sync_message": "",
	}).Error
	if err != nil {
		response.Error(c, err, respstatus.SyncResourceError)
		return
	}

	// 同步云资源
	go func(cloudAccountId string) {
		sync.CloudSyncResource(cloudAccountId)
	}(params.CloudAccountId)

	response.OK(c, nil, "")
}
