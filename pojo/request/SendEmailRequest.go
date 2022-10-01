package request

// SendEmailRequest 定义接收数据的结构体
type SendEmailRequest struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	SystemName string   `form:"SystemName" json:"SystemName" uri:"SystemName" xml:"SystemName" binding:"required"`
	EmailTo    []string `form:"emailTo" json:"emailTo" uri:"emailTo" xml:"emailTo" binding:"emailTo"`
	Subject    string   `form:"subject" json:"subject" uri:"subject" xml:"subject" binding:"required"`
	Content    string   `form:"content" json:"content" uri:"content" xml:"content" binding:"required"`
}
