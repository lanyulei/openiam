package models

type UserGroupRelated struct {
	UserId  int `gorm:"column:user_id;comment:用户ID;uniqueIndex:idx_user_group" json:"user_id" binding:"required"`
	GroupId int `gorm:"column:group_id;comment:用户组ID;uniqueIndex:idx_user_group" json:"group_id" binding:"required"`
}

func (UserGroupRelated) TableName() string {
	return "system_user_group_related"
}
