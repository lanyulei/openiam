package models

import (
	"openiam/common/models"
)

type ApiGroup struct {
	Name   string `gorm:"column:name;type:varchar(200);comment:名称" json:"name"`
	Sort   int    `gorm:"column:sort;comment:排序" json:"sort"`
	Remark string `gorm:"column:remark;type:text;comment:备注" json:"remark"`
	models.BaseModel
}

func (ApiGroup) TableName() string {
	return "system_api_group"
}
