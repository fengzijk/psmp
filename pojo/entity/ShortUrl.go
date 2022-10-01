package entity

import "time"

type ShortURLEntity struct {
	// 主键id
	ID         int64 `gorm:"primarykey;autoIncrement:false"`
	Md5Code    string
	LongParam  string
	ShortParam string
	BizType    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// TableName 表名
func (ShortURLEntity) TableName() string { return "short_param_url" }
