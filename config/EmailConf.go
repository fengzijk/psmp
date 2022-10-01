package config

import (
	"fmt"
	"github.com/spf13/viper"
	"short-url/pojo/model"
)

var EmailConf model.EmailConfModel

func InitEmail() {
	EmailConf := model.EmailConfModel{
		UserName: viper.GetString("redis.address"),
		Password: viper.GetString("redis.password"), // no password set
		Host:     viper.GetString("redis.database"), // use default DB
	}

	fmt.Println(EmailConf)
}
