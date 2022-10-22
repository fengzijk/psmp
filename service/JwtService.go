package service

import (
	"go-psmp/config"
	"go-psmp/utils/date"
	"go-psmp/utils/redis"
)

type JwtService struct {
}

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) bool {
	// 此处过期时间等于jwt过期时间
	dr, err := date.ParseDuration(config.JwtConf.ExpiresTime)
	if err != nil {
		return false
	}
	timer := dr
	flag := redis.Set(userName, jwt, int(timer.Seconds()))
	return flag
}

func (jwtService *JwtService) GetRedisJWT(userName string) string {

	redisJwt := redis.Get(userName)
	return redisJwt
}
