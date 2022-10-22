package config

import (
	"github.com/spf13/viper"
	"go-psmp/pojo/model"
	"golang.org/x/sync/singleflight"
)

var (
	JwtConf            model.JWTConfigModel
	ConcurrencyControl = &singleflight.Group{}
)

func InitBaseConfig() {

	x := model.JWTConfigModel{
		SigningKey:  viper.GetString("jwt.signing-key"),
		ExpiresTime: viper.GetString("jwt.expires-time"),
		BufferTime:  viper.GetString("jwt.buffer-time"),
		Issuer:      viper.GetString("jwt.issuer"), // no password set
	}
	JwtConf = x

}
