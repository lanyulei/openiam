package models

import (
	"openops/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelGroup struct {
	Name  string `json:"name" gorm:"column:name;type:varchar(128);not null;comment:名称" binding:"required"`
	Desc  string `json:"desc" gorm:"column:desc;type:varchar(512);not null;comment:描述"`
	Order int    `json:"order" gorm:"column:order;type:int;not null;default:1;comment:排序"`
	models.BaseModel
}

func (m *ModelGroup) TableName() string {
	return "resource_model_group"
}

func (m *ModelGroup) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.New().String()
	return
}
