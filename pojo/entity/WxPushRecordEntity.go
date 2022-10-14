package entity

import "time"

type WxPushRecordEntity struct {
	// 主键id
	ID            int64  `gorm:"primarykey;autoIncrement:false;comment:主键ID"`
	Md5Code       string `gorm:"type:varchar(20); NOT NULL ; default:''; uniqueIndex: idx_md5,sort:desc; comment:MD5值 唯一"`
	ToUser        string `gorm:"type:varchar(64); NOT NULL; default :''; comment:接收用户账户"`
	ToPartyId     string `gorm:"type:varchar(512) ; NOT NULL; default :''; comment:接收部门"`
	AgentId       string `gorm:"type:varchar(512) ; NOT NULL; default :''; comment:应用ID"`
	Body          string `gorm:"type:varchar(1000) ; NOT NULL; default :''; comment:内容正文 "`
	SendStatus    string `gorm:"type:varchar(8) ; NOT NULL; default :''; comment:发送状态，SUCCESS：发送成功 WAIT：待发送 ERROR：发送异常"`
	ErrorMsg      string `gorm:"type:varchar(1000) ; NOT NULL; default :''; comment:类型 URL-重定向地址 Param-参数 "`
	SendFailCount int    `gorm:"type:int(3); NOT NULL ; default:0;  comment:发送失败次数"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// TableName 表名
func (WxPushRecordEntity) TableName() string { return "wx_push_record" }
