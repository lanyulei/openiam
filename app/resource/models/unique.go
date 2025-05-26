package models

import (
	"openops/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Unique struct {
	Title   string     `json:"title" gorm:"column:title;type:varchar(128);not null;comment:标题"`
	Type    UniqueType `json:"type" gorm:"column:type;type:varchar(128);not null;comment:类型"`
	FieldId string     `json:"field_id" gorm:"column:field_id;type:varchar(128);not null;comment:字段ID"`
	ModelId string     `json:"model_id" gorm:"column:model_id;type:varchar(128);not null;comment:模型ID"`
	Desc    string     `json:"desc" gorm:"column:desc;type:varchar(512);not null;comment:描述"`
	models.BaseModel
}

func (m *Unique) TableName() string {
	return "resource_model_unique"
}

func (m *Unique) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.New().String()
	return
}

type UniqueType string

const (
	// 单独唯一、联合唯一
	UniqueTypeSingle UniqueType = "single"
	UniqueTypeJoint  UniqueType = "joint"
)
