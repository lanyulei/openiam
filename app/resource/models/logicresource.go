package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"openops/common/models"
)

/*
@Author : lanyulei
@Desc :
*/

type LogicResource struct {
	Label   string `gorm:"column:label;type:varchar(128);comment:标签" json:"label"`
	Value   string `gorm:"column:value;type:varchar(128);comment:值" json:"value"`
	Remarks string `gorm:"column:remarks;type:varchar(1024);comment:备注" json:"remarks"`
	models.BaseModel
}

func (l *LogicResource) TableName() string {
	return "resource_logic_resource"
}

func (l *LogicResource) BeforeCreate(tx *gorm.DB) (err error) {
	l.Id = uuid.New().String()
	return
}
