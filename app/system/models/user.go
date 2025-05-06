package models

import (
	"openiam/common/models"
)

type User struct {
	Username   string   `gorm:"column:username;type:varchar(100);comment:用户名" json:"username" binding:"required"`
	Password   string   `gorm:"column:password;type:varchar(100);comment:密码" json:"-"`
	Nickname   string   `gorm:"column:nickname;type:varchar(100);comment:姓名" json:"nickname" binding:"required"`
	Avatar     string   `gorm:"column:avatar;type:varchar(256);comment:头像" json:"avatar"`
	Department int      `gorm:"column:department;comment:部门" json:"department"`
	Tel        string   `gorm:"column:tel;type:varchar(100);comment:手机号" json:"tel"`
	Email      string   `gorm:"column:email;type:varchar(100);comment:邮箱" json:"email" binding:"required,email"`
	Status     bool     `gorm:"column:status;type:boolean;comment:状态" json:"status"`
	IsAdmin    bool     `gorm:"column:is_admin;type:boolean;comment:是否管理员" json:"is_admin"`
	Remark     string   `gorm:"column:remark;type:text;comment:备注" json:"remark"`
	Role       []string `gorm:"-" json:"role"`
	models.BaseModel
}

func (User) TableName() string {
	return "system_user"
}

type UserRequest struct {
	User
	Password string `json:"password"`
}

type UserDetails struct {
	User
	DepartmentInfo Department `gorm:"-" json:"department_info"`
}
