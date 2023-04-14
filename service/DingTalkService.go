package service

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-psmp/config"
	"io"
	"net/http"
	"time"
)

type DingTalkService struct {
}

// SendMessage Function to send message
//
//goland:noinspection GoUnhandledErrorResult
func (t *DingTalkService) SendMessage(s string, at ...string) {
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": s + "\n",
		},
	}

	if config.DingTalkConf.EnableAt {
		if config.DingTalkConf.AtAll {
			if len(at) > 0 {
				config.Log.Error("the parameter AtAll is true,but the at parameter of SendMessage is not empty")
				return
			}
			msg["at"] = map[string]interface{}{
				"isAtAll": config.DingTalkConf.AtAll,
			}
		} else {
			msg["at"] = map[string]interface{}{
				"atMobiles": at,
				"isAtAll":   config.DingTalkConf.AtAll,
			}
		}
	} else {
		if len(at) > 0 {
			config.Log.Error("the parameter \"EnableAt\" is \"false\", but the \"at\" parameter of SendMessage is not empty")
			return
		}
	}
	b, err := json.Marshal(msg)
	if err != nil {
		return
	}
	resp, err := http.Post(t.getURL(), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return
	}

	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Printf(string(body))
	if err != nil {
		return
	}

}

func (t *DingTalkService) hmacSha256(stringToSign string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (t *DingTalkService) getURL() string {
	wh := "https://oapi.dingtalk.com/robot/send?access_token=" + config.DingTalkConf.AccessToken
	timestamp := time.Now().UnixNano() / 1e6
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, config.DingTalkConf.Secret)
	sign := t.hmacSha256(stringToSign, config.DingTalkConf.Secret)
	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", wh, timestamp, sign)
	return url
}
