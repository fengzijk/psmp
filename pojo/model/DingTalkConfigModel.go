package model

type DingTalkConfigModel struct {
	AccessToken string `json:"accessToken"`
	Secret      string `json:"secret"`
	EnableAt    bool   `json:"enableAt"`
	AtAll       bool   `json:"atAll"`
	AtMobile    string `json:"atMobile"`
}
