package entity

import (
	"go-psmp/pojo/model"
)

type UserInfoEntity struct {
	ID        int64  `gorm:"primarykey;autoIncrement:false;comment:主键ID"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	CreatedAt *model.LocalTime
	UpdatedAt *model.LocalTime
}

// TableName 表名
func (UserInfoEntity) TableName() string { return "user_info" }
