package redis

import (
	"encoding/json"
	_ "github.com/go-redis/redis/v8"
	"go-psmp/config"
	"log"
	"sync"

	"time"
)

var mutex sync.Mutex

const (
	redisExpired = time.Duration(0)
)

func Get(key string) string {

	result, err := config.RedisDb.Get(config.Ctx, key).Result()

	if err != nil {
		log.Println(err)
	}

	return result

}

func Set(key string, value string, seconds int) bool {

	ex := redisExpired

	if seconds > 0 {
		ex = time.Duration(seconds) * time.Second
	}

	if len(key) != 0 && "" != key && value != "" {

		set := config.RedisDb.Set(config.Ctx, key, value, ex)
		result, err := set.Result()
		if err != nil {
			log.Println(err)
		}
		return result == "OK"
	}
	return false
}

func SetObj(key string, value interface{}, seconds int) bool {
	ex := redisExpired

	if seconds > 0 {
		ex = time.Duration(seconds) * time.Second
	}

	doctorJson, _ := json.Marshal(value)
	set := config.RedisDb.Set(config.Ctx, key, doctorJson, ex)

	result, err := set.Result()
	if err != nil {
		log.Println(err)
	}
	return result == "OK"

}

func GetObj(key string, value interface{}) {

	var result = value
	cmr, err := config.RedisDb.Get(config.Ctx, key).Result()

	if err != nil {
		log.Println(err)
	}

	if cmr == "" {
		cmr = "{}"
	}

	err = json.Unmarshal([]byte(cmr), &result)
	if err != nil {
		log.Println(err)
	}

}

func Lock(key string, expired int) bool {

	ex := redisExpired

	if expired > 0 {
		ex = time.Duration(expired) * time.Second
	}

	mutex.Lock()
	defer mutex.Unlock()
	result, err := config.RedisDb.SetNX(config.Ctx, key, 1, ex).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return result
}

func UnLock(key string) int64 {
	nums, err := config.RedisDb.Del(config.Ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return nums
}
