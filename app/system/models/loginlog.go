package models

import (
	"openiam/common/models"
)

type LoginLog struct {
	Username string `gorm:"column:username;type:varchar(100);comment:用户名" json:"username"`
	Status   string `gorm:"column:status;type:varchar(100);comment:登陆状态" json:"status"`
	IP       string `gorm:"column:ip;type:varchar(50);comment:IP地址" json:"ip"`
	Browser  string `gorm:"column:browser;type:varchar(200);comment:浏览器" json:"browser"`
	System   string `gorm:"column:system;type:varchar(200);comment:系统" json:"system"`
	Remark   string `gorm:"column:remark;type:text;comment:备注" json:"remark"`
	models.BaseModel
}

func (LoginLog) TableName() string {
	return "system_login_log"
}
