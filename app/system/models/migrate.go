package models

import (
	"openops/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

/*
  @Author : lanyulei
  @Desc :
*/

type Migrate struct {
	Name   string `gorm:"column:name;type:varchar(256);comment:迁移文件名称" json:"name"`
	Status string `gorm:"column:status;type:varchar(45);comment:迁移状态" json:"status"`
	Result string `gorm:"column:result;type:text;comment:迁移结果" json:"result"`
	models.BaseModel
}

func (m *Migrate) TableName() string {
	return "system_migrate"
}

func (m *Migrate) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.New().String()
	return
}
