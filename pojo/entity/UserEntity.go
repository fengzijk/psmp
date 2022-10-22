package entity

import (
	"go-psmp/pojo/model"
)

type UserInfoEntity struct {
	ID           int64  `gorm:"primarykey;autoIncrement:false;comment:主键ID"`
	Username     string `json:"username"`
	Nickname     string `json:"nickname"`
	Password     string `json:"password"`
	PasswordSalt string `json:"passwordSalt"`
	AuthorityId  int    `json:"authorityId"`
	CreatedAt    *model.LocalTime
	UpdatedAt    *model.LocalTime
}

// TableName 表名
func (UserInfoEntity) TableName() string { return "user_info" }
