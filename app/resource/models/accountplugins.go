package models

import (
	"openops/common/models"
)

/*
  @Author : lanyulei
  @Desc :
*/

type AccountPlugins struct {
	AccountId   string `gorm:"column:account_id;type:varchar(64);not null;comment:账号ID" json:"account_id"`
	PluginsName string `gorm:"column:plugins_name;type:varchar(128);not null;comment:插件名称" json:"plugins_name"`
	Remarks     string `gorm:"column:remarks;type:text;comment:备注" json:"remarks"`
	models.BaseModel
}

func (AccountPlugins) TableName() string {
	return "resource_account_plugins"
}
