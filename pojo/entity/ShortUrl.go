package entity

import "time"

type ShortURL struct {
	ID          int64 `gorm:"primarykey;autoIncrement:false"`
	Md5Code     string
	LongParam   string
	ShortParam  string
	RedirectUrl string
	BizType     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// TableName 表名
func (ShortURL) TableName() string { return "short_param_url" }
