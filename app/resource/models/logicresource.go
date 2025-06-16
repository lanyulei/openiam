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
	Title   string `gorm:"column:title;type:varchar(128);comment:标题" json:"title"` // 云主机
	Name    string `gorm:"column:name;type:varchar(128);comment:名称" json:"name"`   // Host
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
