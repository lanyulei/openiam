package models

import (
	"encoding/json"
	"openops/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Data struct {
	ModelId string          `json:"model_id" gorm:"column:model_id;type:varchar(128);not null;comment:模型ID" binding:"required"`
	Data    json.RawMessage `json:"data" gorm:"column:data;type:jsonb;not null;comment:数据"`
	Status  DataStatus      `json:"status" gorm:"column:status;type:varchar(128);not null;comment:状态"`
	models.BaseModel
}

func (d *Data) TableName() string {
	return "resource_data"
}

func (d *Data) BeforeCreate(tx *gorm.DB) (err error) {
	d.Id = uuid.New().String()
	return
}

type DataStatus string

// 默认状态、使用中、空闲中、故障中、维护中、已停止、待回收、已回收
const (
	DataStatusDefault  DataStatus = "default"
	DataStatusUsing    DataStatus = "using"
	DataStatusIdle     DataStatus = "idle"
	DataStatusFault    DataStatus = "fault"
	DataStatusMaintain DataStatus = "maintain"
	DataStatusStopped  DataStatus = "stopped"
	DataStatusRecycle  DataStatus = "recycle"
	DataStatusDeleted  DataStatus = "deleted"
)
