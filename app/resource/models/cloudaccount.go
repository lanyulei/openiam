package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"openops/common/models"
	"openops/pkg/cloud/types"
	"time"
)

/*
  @Author : lanyulei
  @Desc :
*/

type SyncStatusType string

const (
	SyncStatusSuccess SyncStatusType = "success"
	SyncStatusRunning SyncStatusType = "running"
	SyncStatusFailed  SyncStatusType = "failed"
)

type CloudAccount struct {
	Provider     types.CloudName `gorm:"column:provider;type:varchar(45);comment:云服务商;not null;" json:"provider" binding:"required"`
	Name         string          `gorm:"column:name;type:varchar(128);comment:名称;not null;" json:"name" binding:"required"`
	AccessKey    string          `gorm:"column:access_key;type:varchar(128);comment:AccessKey;not null;" json:"access_key"`
	SecretKey    string          `gorm:"column:secret_key;type:varchar(128);comment:SecretKey;not null;" json:"secret_key"`
	Status       bool            `gorm:"column:status;type:boolean;comment:状态" json:"status"`                 // true: 启用，false: 禁用
	SyncStatus   SyncStatusType  `gorm:"column:sync_status;type:varchar(45);comment:同步状态" json:"sync_status"` // success、running、failed
	SyncMessage  string          `gorm:"column:sync_message;type:text;comment:同步信息" json:"sync_message"`
	Available    bool            `gorm:"column:available;type:boolean;comment:是否可用" json:"available"` // 账号是否可连通，true: 可用，false: 不可用
	PluginName   string          `gorm:"column:plugin_name;type:varchar(128);comment:插件名称" json:"plugin_name"`
	LastSyncTime time.Time       `gorm:"column:last_sync_time;comment:最后同步时间" json:"last_sync_time"`
	Type         string          `gorm:"column:type;type:varchar(45);comment:账号类型" json:"type"` // common 通用
	Remarks      string          `gorm:"column:remarks;type:text;comment:备注" json:"remarks"`
	ProviderName string          `gorm:"-" json:"provider_name"`
	models.BaseModel
}

func (c *CloudAccount) TableName() string {
	return "resource_cloud_account"
}

func (c *CloudAccount) BeforeCreate(tx *gorm.DB) (err error) {
	c.Id = uuid.New().String()
	return
}
