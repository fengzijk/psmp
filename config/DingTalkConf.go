package config

import (
	"github.com/spf13/viper"
	"go-psmp/pojo/model"
)

var DingTalkConf model.DingTalkConfigModel

func InitDingTalkConf() {
	x := model.DingTalkConfigModel{
		AccessToken: viper.GetString("dingTalk.accessToken"),
		Secret:      viper.GetString("dingTalk.secret"), // no password set
		AtAll:       viper.GetBool("dingTalk.atAll"),
		EnableAt:    viper.GetBool("dingTalk.enableAt"),
		AtMobile:    viper.GetString("dingTalk.atMobile"),
	}
	DingTalkConf = x

}
