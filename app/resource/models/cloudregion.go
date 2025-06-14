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

type CloudRegion struct {
	CloudAccountId string `gorm:"column:cloud_account_id;type:varchar(128);comment:云账号 ID;not null;" json:"cloud_account_id" binding:"required"`
	RegionId       string `gorm:"column:region_id;type:varchar(128);comment:RegionId;not null;" json:"region_id" binding:"required"`
	Name           string `gorm:"column:name;type:varchar(128);comment:名称;not null;" json:"name" binding:"required"`
	models.BaseModel
}

func (c *CloudRegion) TableName() string {
	return "resource_cloud_region"
}

func (c *CloudRegion) BeforeCreate(tx *gorm.DB) (err error) {
	c.Id = uuid.New().String()
	return
}
