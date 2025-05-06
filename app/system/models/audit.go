package models

import (
	"encoding/json"
	"openiam/common/models"
)

/*
  @Author : lanyulei
  @Desc :
*/

type Audit struct {
	Username string          `gorm:"column:username;type:varchar(100);comment:用户名" json:"username"`
	Path     string          `gorm:"column:path;type:varchar(256);comment:路由" json:"path"`
	Method   string          `gorm:"column:method;type:varchar(50);comment:请求方法" json:"method"`
	Browser  string          `gorm:"column:browser;type:varchar(100);comment:浏览器" json:"browser"`
	IP       string          `gorm:"column:ip;type:varchar(50);comment:ip" json:"ip"`
	System   string          `gorm:"column:system;type:varchar(100);comment:系统" json:"system"`
	Query    json.RawMessage `gorm:"column:query;type:json;comment:请求参数" json:"query"`
	Data     json.RawMessage `gorm:"column:data;type:json;comment:请求体" json:"data"`
	models.BaseModel
}

func (Audit) TableName() string {
	return "system_audit"
}
