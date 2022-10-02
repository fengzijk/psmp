package config

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"github.com/spf13/viper"
	"go-psmp/pojo/model"
	"strings"
)

var EmailConf model.EmailConfModel

var Message *gomail.Message

func InitEmail() {
	x := model.EmailConfModel{
		User:     viper.GetString("email.user"),
		Password: viper.GetString("email.password"), // no password set
		Host:     viper.GetString("email.host"),
		Port:     viper.GetInt("email.port"), // use default DB
	}
	EmailConf = x

	//fmt.Println(EmailConf.User)

	toUser := viper.GetString("email.toUser")
	CCUser := viper.GetString("email.ccUser")

	var towers []string

	//serverHost = ep.ServerHost
	//serverPort = ep.ServerPort
	fromEmail := viper.GetString("email.user")
	fromPasswd := viper.GetString("email.password")
	fmt.Println(fromPasswd)
	Message = gomail.NewMessage()

	if len(toUser) == 0 {
		return
	}

	for _, tmp := range strings.Split(toUser, ";") {
		towers = append(towers, strings.TrimSpace(tmp))
	}

	// 收件人可以有多个，故用此方式
	Message.SetHeader("To", towers...)

	//抄送列表
	if len(CCUser) != 0 {
		for _, tmp := range strings.Split(CCUser, ";") {
			towers = append(towers, strings.TrimSpace(tmp))
		}
		Message.SetHeader("Cc", towers...)
	}

	// 发件人
	// 第三个参数为发件人别名，如"李大锤"，可以为空（此时则为邮箱名称）
	Message.SetAddressHeader("From", fromEmail, "告警平台")
}
