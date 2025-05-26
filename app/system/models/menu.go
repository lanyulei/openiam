package models

import (
	"openops/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Menu 菜单表 postgres
type Menu struct {
	Name        string `json:"name" gorm:"column:name;type:varchar(128);not null;comment:名称"`
	Path        string `json:"path" gorm:"column:path;type:varchar(256);not null;comment:路径"`
	Component   string `json:"component" gorm:"column:component;type:varchar(256);not null;comment:组件"`
	ParentId    string `json:"parent_id" gorm:"column:parent_id;type:varchar(128);not null;comment:父菜单 ID"`
	Redirect    string `json:"redirect" gorm:"column:redirect;type:varchar(256);comment:重定向"`
	Title       string `json:"title" gorm:"column:title;type:varchar(128);comment:标题"`
	Hyperlink   string `json:"hyperlink" gorm:"column:hyperlink;type:varchar(512);comment:超链接"`
	IsHide      bool   `json:"is_hide" gorm:"column:is_hide;type:boolean;not null;default:false;comment:是否隐藏"`
	IsKeepAlive bool   `json:"is_keep_alive" gorm:"column:is_keep_alive;type:boolean;not null;default:true;comment:是否缓存"`
	IsAffix     bool   `json:"is_affix" gorm:"column:is_affix;type:boolean;not null;default:false;comment:是否固定"`
	IsIframe    bool   `json:"is_iframe" gorm:"column:is_iframe;type:boolean;not null;default:false;comment:是否内嵌"`
	IsVerify    bool   `json:"is_verify" gorm:"column:is_verify;type:boolean;not null;default:true;comment:是否验证"`
	Sort        int    `json:"sort" gorm:"column:sort;type:integer;not null;default:0;comment:排序"`
	Icon        string `json:"icon" gorm:"column:icon;type:varchar(128);comment:图标"`
	models.BaseModel
}

func (m *Menu) TableName() string {
	return "system_menu"
}

func (m *Menu) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = uuid.New().String()
	return
}

type MenuMeta struct {
	Title       string `json:"title"`
	Hyperlink   string `json:"hyperlink"`
	IsHide      bool   `json:"isHide"`
	IsKeepAlive bool   `json:"isKeepAlive"`
	IsAffix     bool   `json:"isAffix"`
	IsIframe    bool   `json:"isIframe"`
	Icon        string `json:"icon"`
}

type MenuTree struct {
	Id        string      `json:"id"`
	Name      string      `json:"name"`
	Path      string      `json:"path"`
	Component string      `json:"component"`
	ParentId  string      `json:"parentId"`
	Redirect  string      `json:"redirect"`
	Meta      MenuMeta    `json:"meta"`
	Children  []*MenuTree `json:"children"`
}
