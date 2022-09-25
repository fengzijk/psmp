package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var Ctx = context.Background()
var RedisDb *redis.Client

func InitRedisDb() {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.address"),
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.database"),    // use default DB
	})

}
