package request

// SendEmailRequest 定义接收数据的结构体
type SendEmailRequest struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	SystemName string   `form:"systemName" json:"systemName"`
	EmailTo    []string `form:"emailTo" json:"emailTo" `
	Subject    string   `form:"subject" json:"subject" `
	Content    string   `form:"content" json:"content" `
}
