package models

import (
	"openiam/common/models"
)

type AppGroup struct {
	Name    string `gorm:"column:name;type:varchar(128);comment:应用组名称" json:"name" binding:"required"`
	Sort    int    `gorm:"column:sort;default:1;comment:排序" json:"sort"`
	Remarks string `gorm:"column:remarks;type:varchar(512);comment:备注" json:"remarks"`
	models.BaseModel
}

func (AppGroup) TableName() string {
	return "system_app_group"
}
