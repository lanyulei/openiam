package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id        string         `gorm:"column:id;type:varchar(128);primaryKey;not null;comment:主键ID" json:"id" form:"id"`
	CreatedAt time.Time      `gorm:"column:create_time" json:"create_time" form:"create_time"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"update_time" form:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"column:delete_time" sql:"index" json:"-"`
}

type DBType string

const (
	DBTypeMySQL    DBType = "mysql"
	DBTypePostgres DBType = "postgres"
)
