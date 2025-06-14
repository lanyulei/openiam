package sync

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"solar/app/cmdb/models"
	"solar/pkg/plugin"
	"solar/pkg/tools/comparemaps"
	"time"

	"github.com/lanyulei/toolkit/db"
	"github.com/lanyulei/toolkit/logger"
	"github.com/spf13/viper"
)

/*
  @Author : lanyulei
  @Desc :
*/

func CloudSyncResource(cloudAccountId int) {
	defer func() {
		err := recover()
		if err != nil {
			err = db.Orm().Model(&models.CloudAccount{}).Where("id = ?", cloudAccountId).Update("sync_status", models.SyncStatusFailed).Error
			if err != nil {
				logger.Errorf("recover: update cloud account sync status failed with error: %v", err)
			}
		}
	}()

	var (
		err               error
		cloudInfo         models.CloudAccount
		regions           []string
		cloudModels       []models.CloudModels
		syncStatus        = models.SyncStatusSuccess
		syncMessage       string
		logicResourceMap  = make(map[int]string)
		logicHandleMap    = make(map[int]string)
		logicResourceList []models.LogicResource
		logicHandleList   []models.LogicHandle
	)

	// 获取云账号信息
	err = db.Orm().Model(&models.CloudAccount{}).Where("id = ?", cloudAccountId).First(&cloudInfo).Error
	if err != nil {
		err = fmt.Errorf("get cloud account failed with error: %v", err.Error())
		syncStatus = models.SyncStatusFailed
		goto UpdateSyncStatus
	}

	// 云账号不可用
	if !cloudInfo.Available {
		err = fmt.Errorf("cloud account is not available")
		syncStatus = models.SyncStatusFailed
		goto UpdateSyncStatus
	}

	err = db.Orm().Model(&models.CloudRegion{}).Where("cloud_account_id = ?", cloudAccountId).Pluck("region_id", &regions).Error
	if err != nil {
		err = fmt.Errorf("get cloud region failed with error: %v", err.Error())
		syncStatus = models.SyncStatusFailed
		goto UpdateSyncStatus
	}

	if len(regions) == 0 {
		err = fmt.Errorf("region id is empty")
		syncStatus = models.SyncStatusFailed
		goto UpdateSyncStatus
	}

	// 获取同步设置的资源类型
	err = db.Orm().Model(&models.CloudModels{}).Where("cloud_account_id = ?", cloudAccountId).Find(&cloudModels).Error
	if err != nil {
		err = fmt.Errorf("get cloud models failed with error: %v", err.Error())
		syncStatus = models.SyncStatusFailed
		goto UpdateSyncStatus
	}

	// 获取逻辑资源
	err = db.Orm().Model(&models.LogicResource{}).Find(&logicResourceList).Error
	if err != nil {
		err = fmt.Errorf("get logic resource failed with error: %v", err.Error())
		syncStatus = models.SyncStatusFailed
		goto UpdateSyncStatus
	}

	for _, logicResource := range logicResourceList {
		logicResourceMap[logicResource.Id] = logicResource.Value
	}

	// 获取逻辑处理
	err = db.Orm().Model(&models.LogicHandle{}).Find(&logicHandleList).Error
	if err != nil {
		err = fmt.Errorf("get logic handle failed with error: %v", err.Error())
		syncStatus = models.SyncStatusFailed
		goto UpdateSyncStatus
	}

	for _, logicHandle := range logicHandleList {
		logicHandleMap[logicHandle.Id] = logicHandle.Name
	}

	if len(cloudModels) == 0 {
		err = fmt.Errorf("cloud models is empty")
		syncStatus = models.SyncStatusFailed
		goto UpdateSyncStatus
	}

	for _, region := range regions {
		for _, cloudModel := range cloudModels {
			var (
				resp         []byte
				result       []plugin.Response
				resourceInfo models.Resource
			)

			pluginPath := filepath.Join(viper.GetString("plugin.path"), cloudInfo.PluginName)
			resp, err = plugin.New(pluginPath).Get(context.TODO(), logicResourceMap[cloudModel.LogicResource], region, logicHandleMap[cloudModel.LogicHandle], nil)
			if err != nil {
				err = fmt.Errorf("failed to get cloud resource error: %v", err.Error())
				syncStatus = models.SyncStatusFailed
				goto UpdateSyncStatus
			}

			err = json.Unmarshal(resp, &result)
			if err != nil {
				err = fmt.Errorf("unserialize cloud resource failed with error: %v", err.Error())
				syncStatus = models.SyncStatusFailed
				goto UpdateSyncStatus
			}

			for _, instance := range result {
				err = db.Orm().Model(&models.Resource{}).Where("model_id = ? and unique_id = ?", cloudModel.ModelId, instance.UniqueId).Find(&resourceInfo).Error
				if err != nil {
					err = fmt.Errorf("get resource failed with error: %v", err.Error())
					syncStatus = models.SyncStatusFailed
					goto UpdateSyncStatus
				}

				if resourceInfo.Id != 0 {
					sourceContent := make(map[string]interface{})
					err = json.Unmarshal(resourceInfo.Content, &sourceContent)
					if err != nil {
						err = fmt.Errorf("unserialize resource content failed with error: %v", err.Error())
						syncStatus = models.SyncStatusFailed
						goto UpdateSyncStatus
					}

					targetContent := make(map[string]interface{})
					err = json.Unmarshal(instance.Content, &targetContent)
					if err != nil {
						err = fmt.Errorf("unserialize instance content failed with error: %v", err.Error())
						syncStatus = models.SyncStatusFailed
						goto UpdateSyncStatus
					}

					// 比较资源内容是否一致
					ok := comparemaps.CompareMaps(sourceContent, targetContent)
					if !ok {
						err = db.Orm().Model(&models.Resource{}).Where("model_id = ? and unique_id = ?", cloudModel.ModelId, instance.UniqueId).Updates(map[string]interface{}{
							"content": instance.Content,
						}).Error
						if err != nil {
							err = fmt.Errorf("update resource failed with error: %v", err.Error())
							syncStatus = models.SyncStatusFailed
							goto UpdateSyncStatus
						}
					}
				} else {
					err = db.Orm().Create(&models.Resource{
						ModelId:  cloudModel.ModelId,
						UniqueId: instance.UniqueId,
						Content:  instance.Content,
					}).Error
					if err != nil {
						err = fmt.Errorf("create resource failed with error: %v", err.Error())
						syncStatus = models.SyncStatusFailed
						goto UpdateSyncStatus
					}
				}
			}
		}
	}
UpdateSyncStatus:
	if err != nil {
		logger.Error(err.Error())
		syncMessage = err.Error()
	} else {
		syncMessage = "sync cloud resource success"
	}
	err = db.Orm().Model(&models.CloudAccount{}).Where("id = ?", cloudAccountId).Updates(map[string]interface{}{
		"sync_status":    syncStatus,
		"sync_message":   syncMessage,
		"last_sync_time": time.Now(),
	}).Error
	if err != nil {
		logger.Errorf("update cloud account sync status failed with error: %s", err.Error())
	}
}
