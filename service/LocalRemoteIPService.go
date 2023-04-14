package service

import (
	"fmt"
	"go-psmp/config"
	"go-psmp/utils/redis"
	"io"

	"net/http"
)

const ipCacheShortKey = "remote_ip_cache"

type LocalRemoteIPService struct {
}

var dingTalkService = ServiceGroup.DingTalkService

func (localRemoteIPService *LocalRemoteIPService) SendIPChange() {
	responseClient, errClient := http.Get("https://ipv4.netarm.com/") // 获取外网 IP
	if errClient != nil {
		fmt.Printf("获取外网 IP 失败，请检查网络\n")
		config.Log.Error(errClient.Error())
	}
	// 程序在使用完 response 后必须关闭 response 的主体。
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(responseClient.Body)

	body, _ := io.ReadAll(responseClient.Body)
	clientIP := fmt.Sprintf("%s", string(body))

	redisCacheKey := ipCacheShortKey

	// 缓存中查
	lastIp := redis.Get(redisCacheKey)

	if lastIp != clientIP {
		redis.Set(redisCacheKey, clientIP, 60*11)
		dingTalkService.SendMessage("你的IP由" + lastIp + "\n变化为\n" + clientIP)

	}

}
