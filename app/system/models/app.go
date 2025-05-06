package models

import (
	"openiam/common/models"
)

type App struct {
	Name       string `gorm:"column:name;type:varchar(128);comment:应用名称" json:"name" binding:"required"`
	Icon       string `gorm:"column:icon;type:varchar(128);comment:应用图标" json:"icon" binding:"required"`
	IsExternal bool   `gorm:"column:is_external;type:boolean;comment:是否外部应用" json:"is_external"`
	Link       string `gorm:"column:link;type:varchar(256);comment:应用链接" json:"link" binding:"required"`
	IsBlank    bool   `gorm:"column:is_blank;type:boolean;comment:是否新窗口打开" json:"is_blank"`
	AppGroupId int    `gorm:"column:app_group_id;comment:应用组ID" json:"app_group_id" binding:"required"`
	Sort       int    `gorm:"column:sort;default:1;comment:排序" json:"sort"`
	Remarks    string `gorm:"column:remarks;type:varchar(512);comment:备注" json:"remarks"`
	models.BaseModel
}

func (App) TableName() string {
	return "system_app"
}
