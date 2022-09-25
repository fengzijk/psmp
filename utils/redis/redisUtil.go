package redis

import (
	"encoding/json"
	_ "github.com/go-redis/redis/v8"
	"log"
	"short-url/config"
	"time"
)

func Get(key string) string {

	result, err := config.RedisDb.Get(config.Ctx, key).Result()

	if err != nil {
		log.Fatal(err)
	}

	return result

}

func Set(key string, value string) {
	if len(key) != 0 && "" != key && value != "" {

		set := config.RedisDb.Set(config.Ctx, key, value, time.Duration(-1))
		result, err := set.Result()
		if err != nil {
			log.Println(err)
		}
		log.Println(result)
	}

}
func SetObj(key string, value interface{}) {
	doctorJson, _ := json.Marshal(value)
	set := config.RedisDb.Set(config.Ctx, key, doctorJson, time.Duration(-1))

	result, err := set.Result()
	if err != nil {
		log.Println(err)
	}
	log.Println(result)

}

func GetObj(key string, value interface{}) interface{} {

	var re = value
	result, err := config.RedisDb.Get(config.Ctx, key).Result()

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(result), &re)
	if err != nil {
		return nil
	}
	return re

}
