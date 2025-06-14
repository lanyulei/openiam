package server

import (
	"github.com/lanyulei/toolkit/db"
	"openops/app/resource/models"
)

/*
  @Author : lanyulei
  @Desc :
*/

func initData() (err error) {
	err = db.Orm().Model(&models.CloudAccount{}).Where("sync_status = ?", models.SyncStatusRunning).Update("sync_status", models.SyncStatusFailed).Error
	if err != nil {
		return
	}
	return
}
