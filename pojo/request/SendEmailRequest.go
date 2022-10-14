package request

// SendEmailRequest 定义接收数据的结构体
type SendEmailRequest struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	FromName string `form:"fromName" json:"fromName"`
	ToUser   string `form:"ToUser" json:"toUser" `
	CcUser   string `form:"emailTo" json:"ccUser" `
	Subject  string `form:"subject" json:"subject" `
	Body     string `form:"content" json:"content" `
}

// AlarmRequest 定义接收数据的结构体
type AlarmRequest struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	FromName string `form:"fromName" json:"fromName"`
	Body     string `form:"content" json:"content" `
}
