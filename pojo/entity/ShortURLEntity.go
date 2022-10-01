package entity

import "time"

type ShortURLEntity struct {
	// 主键id
	ID         int64  `gorm:"primarykey;autoIncrement:false;comment:主键ID"`
	Md5Code    string `gorm:"type:varchar(20); NOT NULL ; default:''; uniqueIndex: idx_md5,sort:desc; comment:长参数MD5code"`
	LongParam  string `gorm:"type:varchar(1000); NOT NULL; default :''; comment:原始参数"`
	ShortParam string `gorm:"type:varchar(20) ; NOT NULL; default :''; uniqueIndex: idx_stp,sort:desc; comment:短参数短参数"`
	BizType    string `gorm:"type:varchar(8) ; NOT NULL; default :''; comment:类型 URL-重定向地址 Param-参数 "`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// TableName 表名
func (ShortURLEntity) TableName() string { return "short_param_url" }
