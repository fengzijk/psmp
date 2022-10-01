package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go-psmp/pojo/model"
)

var EmailConf model.EmailConfModel

func InitEmail() {
	x := model.EmailConfModel{
		User:     viper.GetString("email.user"),
		Password: viper.GetString("email.password"), // no password set
		Host:     viper.GetString("email.host"),     // use default DB
	}
	EmailConf = x

	fmt.Println(EmailConf.User)
}
