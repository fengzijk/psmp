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

func Set(key string, value string) bool {
	if len(key) != 0 && "" != key && value != "" {

		set := config.RedisDb.Set(config.Ctx, key, value, time.Duration(-1))
		result, err := set.Result()
		if err != nil {
			log.Println(err)
		}
		return result == "OK"
	}
	return false
}
func SetObj(key string, value interface{}) bool {
	doctorJson, _ := json.Marshal(value)
	set := config.RedisDb.Set(config.Ctx, key, doctorJson, time.Duration(-1))

	result, err := set.Result()
	if err != nil {
		log.Fatal(err)
	}
	return result == "OK"

}

func GetObj(key string, value interface{}) {

	var result = value
	cmr, err := config.RedisDb.Get(config.Ctx, key).Result()

	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal([]byte(cmr), &result)
	if err != nil {
		log.Println(err)
	}

}
