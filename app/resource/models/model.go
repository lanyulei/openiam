package models

import (
	"encoding/json"
	"openops/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	Name      string          `json:"name" gorm:"column:name;type:varchar(128);not null;comment:名称"`
	Icon      json.RawMessage `json:"icon" gorm:"column:icon;type:jsonb;not null;comment:图标"`
	Status    bool            `json:"status" gorm:"column:status;type:boolean;not null;comment:状态"`
	Desc      string          `json:"desc" gorm:"column:desc;type:varchar(512);not null;comment:描述"`
	GroupId   string          `json:"group_id" gorm:"column:group_id;type:varchar(128);not null;comment:分组ID"`
	Order     int             `json:"order" gorm:"column:order;type:int;not null;comment:排序"`
	DataCount int             `json:"data_count" gorm:"-"`
	models.BaseModel
}

func (m *Model) TableName() string {
	return "resource_model"
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.New().String()
	return
}
