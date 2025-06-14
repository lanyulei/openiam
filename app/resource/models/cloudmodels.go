package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"openops/common/models"
)

/*
  @Author : lanyulei
  @Desc :
*/

type CloudModels struct {
	CloudAccountId string `gorm:"column:cloud_account_id;type:varchar(128);comment:云账号ID;not null;" json:"cloud_account_id" binding:"required"`
	ModelId        string `gorm:"column:model_id;type:varchar(128);comment:模型ID;not null;" json:"model_id" binding:"required"`
	LogicResource  string `gorm:"column:logic_resource;type:varchar(128);comment:逻辑资源" json:"logic_resource" binding:"required"` // Host、Disk
	LogicHandle    string `gorm:"column:logic_handle;type:varchar(128);comment:逻辑处理" json:"logic_handle" binding:"required"`     // ResourceList、CreateResource
	models.BaseModel
}

func (c *CloudModels) TableName() string {
	return "resource_cloud_models"
}

func (c *CloudModels) BeforeCreate(tx *gorm.DB) (err error) {
	c.Id = uuid.New().String()
	return
}
