package service

import (
	"encoding/json"
	"go-psmp/mapper"
	"go-psmp/pojo/dto"
	"go-psmp/pojo/entity"
	util "go-psmp/utils/http"
	"go-psmp/utils/short"
	"log"
)

type Result struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
}

type AccessToken struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
	Time      int64  `json:"time"`
}

func getMessageBody(toUser string, toParty string, agentId int, content string) string {
	msg := dto.WxPushMessage{
		ToUser:  toUser,
		ToParty: toParty,
		MsgType: "text",
		AgentId: agentId,
		Safe:    0,
		Text: struct {
			//Subject string `json:"subject"`
			Content string `json:"content"`
		}{Content: content},
	}
	sendMsg, _ := json.Marshal(msg)

	return string(sendMsg)
}

func getAccessToken(corpId, corpSecret string) string {

	getTokenUrl := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + corpId + "&corpsecret=" + corpSecret

	result := util.GetJson(getTokenUrl, "")

	var accessToken AccessToken
	err := json.Unmarshal([]byte(result), &accessToken)
	if err != nil {
		return ""
	}

	return accessToken.Token
}

func sendMessage(accessToken, msg string) string {
	sendUrl := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + accessToken
	postJson := util.PostJson(sendUrl, msg, "")
	var res Result
	err := json.Unmarshal([]byte(postJson), &res)
	if err != nil {
		log.Print(err)
	}
	return res.Errmsg
}

func SendWxPushMessage(corpId, corpSecret, toUser, toParty, messageContent string, agentId int) string {
	token := getAccessToken(corpId, corpSecret)
	var res string
	if token != "" {
		messageBody := getMessageBody(toUser, toParty, agentId, messageContent)
		res = sendMessage(token, messageBody)

	} else {
		return "get token fail"
	}
	return res
}

func SaveWxPushMessage(toUser, toParty, messageContent string, agentId int) bool {

	insert := entity.WxPushRecordEntity{
		Md5Code: short.Get16MD5Encode(toUser + toParty + messageContent + string(rune(agentId))),

		ToUser:        toUser,
		ToPartyId:     toParty,
		AgentId:       agentId,
		Body:          messageContent,
		SendStatus:    "WAIT",
		SendFailCount: 0,
	}

	err := mapper.InsertWxPushRecord(insert)
	if err != nil {
		return false
	}
	return true
}

func UpdateWxPushSendSuccess(ids []int64) {
	mapper.UpdateWxPushSendSuccess(ids)
}

func UpdateWxPushSendFail(failList []entity.WxPushRecordEntity) {
	mapper.UpdateWxPushSendFail(failList)
}
