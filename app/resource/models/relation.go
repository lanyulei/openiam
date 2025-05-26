package models

import (
	"openops/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Relation struct {
	SourceModelId string             `json:"source_model_id" gorm:"column:source_model_id;type:varchar(128);not null;comment:来源模型ID"`
	TargetModelId string             `json:"target_model_id" gorm:"column:target_model_id;type:varchar(128);not null;comment:目标模型ID"`
	Type          RelationType       `json:"type" gorm:"column:type;type:varchar(128);not null;comment:关联类型"`
	Constraint    RelationConstraint `json:"constraint" gorm:"column:constraint;type:varchar(128);not null;comment:模型约束"`
	Desc          string             `json:"desc" gorm:"column:desc;type:varchar(512);not null;comment:描述"`
	models.BaseModel
}

func (m *Relation) TableName() string {
	return "resource_model_relation"
}

func (m *Relation) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.New().String()
	return
}

type RelationType string

const (
	RelationTypeBelong  RelationType = "belong"  // 属于
	RelationTypeGroup   RelationType = "group"   // 组成
	RelationTypeRun     RelationType = "run"     // 运行
	RelationTypeConnect RelationType = "connect" // 上联
	RelationTypeDefault RelationType = "default" // 默认
)

type RelationConstraint string

const (
	OneToOne   RelationConstraint = "oneToOne"
	OneToMany  RelationConstraint = "oneToMany"
	ManyToMany RelationConstraint = "manyToMany"
)
