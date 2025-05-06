package models

import "openiam/common/models"

type RoleMenu struct {
	Role int `gorm:"column:role;comment:角色" json:"role"`
	Menu int `gorm:"column:menu;comment:菜单" json:"menu"`
	Type int `gorm:"column:type;comment:类型" json:"type"` // 1 菜单， 2 按钮
	models.BaseModel
}

func (RoleMenu) TableName() string {
	return "system_role_menu"
}
