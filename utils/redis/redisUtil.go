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
	prefix       = "psmp:"
)

func Get(key string) string {

	result, err := config.RedisDb.Get(config.Ctx, prefix+key).Result()

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

		set := config.RedisDb.Set(config.Ctx, prefix+key, value, ex)
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
	set := config.RedisDb.Set(config.Ctx, prefix+key, doctorJson, ex)

	result, err := set.Result()
	if err != nil {
		log.Println(err)
	}
	return result == "OK"

}

func GetObj(key string, value interface{}) {

	var result = value
	cmr, err := config.RedisDb.Get(config.Ctx, prefix+key).Result()

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

func HGetAll(key string) map[string]string {

	var result map[string]string
	cmr, err := config.RedisDb.HGetAll(config.Ctx, prefix+key).Result()

	if err != nil {
		log.Println(err)
	}

	if cmr != nil {
		result = cmr
	}

	return result
}

func Exists(key string) bool {

	if len(key) != 0 && "" != key {

		set := config.RedisDb.Exists(config.Ctx, prefix+key)
		result, err := set.Result()
		if err != nil {
			log.Println(err)
		}
		return result > 0
	}

	return false
}

func Delete(key string) bool {

	if len(key) != 0 && "" != key {

		set := config.RedisDb.Del(config.Ctx, prefix+key)
		result, err := set.Result()
		if err != nil {
			log.Println(err)
		}
		return result > 0
	}

	return false
}

func SetEx(key string, value interface{}, seconds int) bool {

	ex := redisExpired

	if seconds > 0 {
		ex = time.Duration(seconds) * time.Second
	}

	if len(key) != 0 && "" != key {

		set := config.RedisDb.SetEX(config.Ctx, prefix+key, value, ex)
		result, err := set.Result()
		if err != nil {
			log.Println(err)
		}
		return result == "OK"
	}

	return false
}

func HSet(key string, value interface{}, seconds int) bool {

	ex := redisExpired

	if seconds > 0 {
		ex = time.Duration(seconds) * time.Second
	}

	if len(key) != 0 && "" != key && value != nil {

		set := config.RedisDb.HSet(config.Ctx, prefix+key, value, ex)
		result, err := set.Result()
		if err != nil {
			log.Println(err)
		}
		return result > 0
	}

	return false
}

func Lock(key string, expired int) bool {

	ex := redisExpired

	if expired > 0 {
		ex = time.Duration(expired) * time.Second
	}

	mutex.Lock()
	defer mutex.Unlock()
	result, err := config.RedisDb.SetNX(config.Ctx, prefix+key, 1, ex).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return result
}

func UnLock(key string) int64 {
	nums, err := config.RedisDb.Del(config.Ctx, prefix+key).Result()
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return nums
}
