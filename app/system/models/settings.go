package models

import (
	"encoding/json"
	"openiam/common/models"
)

/*
  @Author : lanyulei
  @Desc :
*/

type Settings struct {
	Key     string          `gorm:"column:key;type:varchar(100);comment:唯一值" json:"key" binding:"required"`
	Content json.RawMessage `gorm:"column:content;type:json;comment:配置详情" json:"content"`
	models.BaseModel
}

func (Settings) TableName() string {
	return "system_settings"
}
