package entity

import "time"

type EmailRecordEntity struct {
	// 主键id
	ID            int64  `gorm:"primarykey;autoIncrement:false;comment:主键ID"`
	Md5Code       string `gorm:"type:varchar(20); NOT NULL ; default:''; uniqueIndex: idx_md5,sort:desc; comment:MD5值 唯一"`
	FromName      string `gorm:"type:varchar(64); NOT NULL; default :''; comment:发件人"`
	EmailTo       string `gorm:"type:varchar(512) ; NOT NULL; default :''; comment:收件人，多个收件人用英文;分隔"`
	EmailCc       string `gorm:"type:varchar(512) ; NOT NULL; default :''; comment:收件人，多个收件人用英文;分隔"`
	Subject       string `gorm:"type:varchar(200) ; NOT NULL; default :''; comment:标题"`
	Body          string `gorm:"type:varchar(1000) ; NOT NULL; default :''; comment:内容正文 "`
	SendStatus    string `gorm:"type:varchar(8) ; NOT NULL; default :''; comment:发送状态，SUCCESS：发送成功 WAIT：待发送 ERROR：发送异常"`
	ErrorMsg      string `gorm:"type:varchar(1000) ; NOT NULL; default :''; comment:类型 URL-重定向地址 Param-参数 "`
	SendFailCount int    `gorm:"type:int(3); NOT NULL ; default:0;  comment:发送失败次数"`
	TemplateFlag  string `gorm:"type:varchar(8) ; NOT NULL; default :'HTML'; comment:HTML NORMAL"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName 表名
func (EmailRecordEntity) TableName() string { return "email_record" }
