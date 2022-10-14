package dto

type WxPushMessage struct {
	ToUser  string `json:"touser"`
	ToParty string `json:"toparty"`
	MsgType string `json:"msgtype"`
	AgentId int    `json:"agentid"`
	Text    struct {
		//Subject string `json:"subject"`
		Content string `json:"content"`
	} `json:"text"`
	Safe int `json:"safe"`
}
