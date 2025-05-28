package models

import (
	"openops/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FieldGroup struct {
	Name    string `json:"name" gorm:"column:name;type:varchar(128);not null;comment:名称"`
	Desc    string `json:"desc" gorm:"column:desc;type:varchar(512);not null;comment:描述"`
	Order   int    `json:"order" gorm:"column:order;type:int;not null;comment:排序"`
	ModelId string `json:"model_id" gorm:"column:model_id;type:varchar(128);not null;comment:模型ID"`
	models.BaseModel
}

func (m *FieldGroup) TableName() string {
	return "resource_field_group"
}

func (m *FieldGroup) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.New().String()
	return
}
