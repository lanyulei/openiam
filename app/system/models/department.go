package models

import "openiam/common/models"

/*
  @Author : lanyulei
  @Desc :
*/

type Department struct {
	Name    string `gorm:"column:name;type:varchar(100);comment:名称" json:"name" binding:"required"`
	Parent  int    `gorm:"column:parent;comment:父级ID" json:"parent"`
	Leader  int    `gorm:"column:leader;comment:负责人" json:"leader"`
	Remarks string `gorm:"column:remarks;type:varchar(1024);comment:备注" json:"remarks"`
	models.BaseModel
}

func (Department) TableName() string {
	return "system_department"
}

type DepartmentTree struct {
	Department
	LeaderUsername string            `json:"leader_username"`
	LeaderNickname string            `json:"leader_nickname"`
	Children       []*DepartmentTree `gorm:"-" json:"children"`
}
