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
	LogicResourceId string `gorm:"column:logic_resource_id;type:varchar(128);comment:逻辑资源ID" json:"logic_resource_id"` // 逻辑资源ID
	Title           string `gorm:"column:title;type:varchar(128);comment:标题" json:"title" binding:"required"`          // 资源列表、创建资源
	Name            string `gorm:"column:name;type:varchar(128);comment:名称" json:"name" binding:"required"`            // ResourceList、CreateResource
	Remarks         string `gorm:"column:remarks;type:varchar(512);comment:备注" json:"remarks"`
	models.BaseModel
}

func (l *LogicHandle) TableName() string {
	return "resource_logic_handle"
}

func (l *LogicHandle) BeforeCreate(tx *gorm.DB) (err error) {
	l.Id = uuid.New().String()
	return
}
