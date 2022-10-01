package request

type ShortContentRequest struct {
	Content string `json:"content"`
	BizType int    `json:"bizType"`
}
