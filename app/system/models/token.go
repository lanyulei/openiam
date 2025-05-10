package models

import "openiam/common/models"

type TokenStatus string

const (
	TokenStatusValid   TokenStatus = "valid"   // 有效
	TokenStatusInvalid TokenStatus = "invalid" // 无效
)

type TokenType string

const (
	AccessToken  TokenType = "access"  // 访问令牌
	RefreshToken TokenType = "refresh" // 刷新令牌
)

type Token struct {
	UserId    string      `json:"user_id" gorm:"column:user_id;type:varchar(128);not null;comment:用户ID"`
	Username  string      `json:"username" gorm:"column:username;type:varchar(128);not null;comment:用户名"`
	JwtId     string      `json:"jwt_id" gorm:"column:jwt_id;type:varchar(128);not null;comment:JWT ID"`
	IssuedAt  int64       `json:"issued_at" gorm:"column:issued_at;type:bigint;not null;comment:签发时间"`
	ExpiresAt int64       `json:"expires_at" gorm:"column:expires_at;type:bigint;not null;comment:过期时间"`
	Status    TokenStatus `json:"status" gorm:"column:status;type:varchar(32);not null;comment:令牌状态"`
	Type      TokenType   `json:"type" gorm:"column:type;type:varchar(32);not null;comment:令牌类型"`
	models.BaseModel
}

func (Token) TableName() string {
	return "system_token"
}
