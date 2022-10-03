package config

import (
	"github.com/go-gomail/gomail"
	"github.com/spf13/viper"
	"go-psmp/pojo/model"
)

var EmailConf model.EmailConfModel

var Message *gomail.Message

func InitEmail() {
	x := model.EmailConfModel{
		User:     viper.GetString("email.user"),
		Password: viper.GetString("email.password"), // no password set
		Host:     viper.GetString("email.host"),
		Port:     viper.GetInt("email.port"),
	}
	EmailConf = x
	Message = gomail.NewMessage()
}
