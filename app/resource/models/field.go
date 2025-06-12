package models

import (
	"encoding/json"
	"openops/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Field struct {
	Key         string          `json:"key" gorm:"column:key;type:varchar(128);not null;comment:字段标识"`
	Name        string          `json:"name" gorm:"column:name;type:varchar(128);not null;comment:名称"`
	GroupId     string          `json:"group_id" gorm:"column:group_id;type:varchar(128);not null;comment:分组ID"`
	Type        FieldType       `json:"type" gorm:"column:type;type:varchar(128);not null;comment:类型"`
	Options     json.RawMessage `json:"options" gorm:"column:options;type:jsonb;not null;comment:选项"`
	IsEdit      bool            `json:"is_edit" gorm:"column:is_edit;type:boolean;not null;default:true;comment:是否可编辑"`
	IsRequired  bool            `json:"is_required" gorm:"column:is_required;type:boolean;default:false;not null;comment:是否必填"`
	IsList      bool            `json:"is_list" gorm:"column:is_list;type:boolean;not null;default:false;comment:是否列表展示"` // FieldTypeTable 时，无法在列表展示
	Placeholder string          `json:"placeholder" gorm:"column:placeholder;type:varchar(256);not null;comment:占位符"`
	Desc        string          `json:"desc" gorm:"column:desc;type:varchar(512);not null;comment:描述"`
	Order       int             `json:"order" gorm:"column:order;type:int;not null;default:0;comment:排序"`
	Span        int             `json:"span" gorm:"column:span;type:int;not null;default:12;comment:栅格数"`
	ModelId     string          `json:"model_id" gorm:"column:model_id;type:varchar(128);not null;comment:模型ID"`
	models.BaseModel
}

func (m *Field) TableName() string {
	return "resource_field"
}

func (m *Field) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.New().String()
	return
}

type FieldType string

const (
	FieldTypeShortString FieldType = "shortString"
	FieldTypeNumber      FieldType = "number"
	FieldTypeFloat       FieldType = "float"
	FieldTypeEnum        FieldType = "enum"
	FieldTypeEnumMulti   FieldType = "enumMulti"
	FieldTypeDate        FieldType = "date"
	FieldTypeTime        FieldType = "time"
	FieldTypeDateTime    FieldType = "dateTime"
	FieldTypeLongString  FieldType = "longString"
	FieldTypeUser        FieldType = "user"
	FieldTypeTimeZone    FieldType = "timeZone"
	FieldTypeBoolean     FieldType = "boolean"
	FieldTypeList        FieldType = "list"
	FieldTypeTable       FieldType = "table"
)

// LabelValue 定义一个结构体来表示标签和值的映射
type LabelValue struct {
	Label string
	Value FieldType
}

// 定义一个基础的 FieldType 映射列表，包含所有类型
var baseFieldTypeMap = []LabelValue{
	{"短字符串", FieldTypeShortString},
	{"数字", FieldTypeNumber},
	{"浮点数", FieldTypeFloat},
	{"枚举", FieldTypeEnum},
	{"枚举多选", FieldTypeEnumMulti},
	{"日期", FieldTypeDate},
	{"时间", FieldTypeTime},
	{"日期时间", FieldTypeDateTime},
	{"长字符串", FieldTypeLongString},
	{"用户", FieldTypeUser},
	{"时区", FieldTypeTimeZone},
	{"布尔值", FieldTypeBoolean},
	{"列表", FieldTypeList},
	{"表格", FieldTypeTable},
}

// TableOptionTypeList 从基础列表中筛选出表格支持的类型
var TableOptionTypeList = func() []LabelValue {
	var result []LabelValue
	allowedTypes := map[FieldType]bool{
		FieldTypeShortString: true,
		FieldTypeNumber:      true,
		FieldTypeFloat:       true,
		FieldTypeEnum:        true,
		FieldTypeLongString:  true,
		FieldTypeBoolean:     true,
	}
	for _, item := range baseFieldTypeMap {
		if allowedTypes[item.Value] {
			result = append(result, item)
		}
	}
	return result
}()

// FieldTypeValueList 直接使用基础列表
var FieldTypeValueList = baseFieldTypeMap

// StringOptions 短字符串及长字符配置，正则表达式、默认值
type StringOptions struct {
	Regexp  string `json:"regexp"`
	Default string `json:"default"`
}

// NumberOptions 数字配置，最小值、最大值、默认值
type NumberOptions struct {
	Min     int `json:"min"`
	Max     int `json:"max"`
	Default int `json:"default"`
}

// FloatOptions 浮点数配置，最小值、最大值、默认值
type FloatOptions struct {
	Min     float64 `json:"min"`
	Max     float64 `json:"max"`
	Default float64 `json:"default"`
}

// EnumOptions 枚举配置，选项
type EnumOptions struct {
	Options []struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"options"`
	Default string `json:"default"`
}

// EnumMultiOptions 多选枚举配置，选项
type EnumMultiOptions struct {
	Options []struct {
		ID    string `json:"id"`
		Value string `json:"value"`
	} `json:"options"`
	Default []string `json:"default"`
}

// DefaultOptions 用户、日期、时间、日期时间、时区的配置，默认值
type DefaultOptions struct {
	Default string `json:"default"`
}

// BooleanOptions 布尔值配置，默认值
type BooleanOptions struct {
	Default bool `json:"default"`
}

// ListOptions 列表配置，默认值
type ListOptions struct {
	Options []string `json:"options"`
	Default string   `json:"default"`
}

// TableOptions 表格配置，默认值
type TableOptions struct {
	Columns []struct {
		Label string `json:"label"`
		Value string `json:"value"`
		// Type       string `json:"type"` // 支持 ShortString、Number、Float、Enum、EnumMulti、Boolean
		// IsEdit     bool   `json:"is_edit"`
		// IsRequired bool   `json:"is_required"`
		// Regexp     string `json:"regexp"`
		// Default    string `json:"default"`
	} `json:"columns"`
	Default []map[string]interface{} `json:"default"`
}
