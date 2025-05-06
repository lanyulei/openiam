package models

import (
	"openiam/common/models"
)

type UserGroup struct {
	Name    string `gorm:"column:name;type:varchar(100);comment:用户组名称" json:"name" binding:"required"`
	App     string `gorm:"column:app;type:varchar(100);comment:应用名称" json:"app" binding:"required"`
	Remarks string `gorm:"column:remarks;type:varchar(1024);comment:备注" json:"remarks"`
	models.BaseModel
}

func (UserGroup) TableName() string {
	return "system_user_group"
}
