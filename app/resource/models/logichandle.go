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

type LogicHandle struct {
	LogicResource string `gorm:"column:logic_resource;type:varchar(128);comment:逻辑资源" json:"logic_resource" binding:"required"` // Host、Disk
	Name          string `gorm:"column:name;type:varchar(128);comment:逻辑处理名称" json:"name"`                                      // ResourceList、CreateResource
	Remarks       string `gorm:"column:remarks;type:varchar(512);comment:备注" json:"remarks"`
	models.BaseModel
}

func (l *LogicHandle) TableName() string {
	return "resource_logic_handle"
}

func (l *LogicHandle) BeforeCreate(tx *gorm.DB) (err error) {
	l.Id = uuid.New().String()
	return
}
