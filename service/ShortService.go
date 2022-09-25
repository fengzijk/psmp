package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"short-url/mapper"
	"short-url/pojo/entity"
	"short-url/utils/redis"
	shortUtil "short-url/utils/short"
)

type Response gin.H

const (
	path     = "st"
	cacheKey = "short:"
)

// CreateShort 生成短连接
func CreateShort(param string, bizType string) string {

	var urlEntity entity.ShortURL

	md5Code := shortUtil.Get16MD5Encode(param)
	shortParam := shortUtil.GetShortParam(param)
	redisCacheKey := cacheKey + md5Code

	// 缓存中查
	redis.GetObj(redisCacheKey, &urlEntity)

	// 数据库查询
	if urlEntity.RedirectUrl == "" || urlEntity.BizType == "url" {
		urlEntity = mapper.SelectShortUrlInfoByMd5Code(md5Code)
	}

	// 存在返回结果
	if urlEntity.RedirectUrl != "" && urlEntity.BizType == "url" {
		redis.SetObj(redisCacheKey, urlEntity)
		return viper.GetString("short.prefix") + "/" + urlEntity.RedirectUrl
	}

	// 新增
	redirectUrl := fmt.Sprintf("%s/%s", path, shortParam)
	shortInfo := entity.ShortURL{
		ShortParam:  shortParam,
		RedirectUrl: redirectUrl,
		LongParam:   param,
		Md5Code:     md5Code,
		BizType:     bizType,
	}
	err := mapper.InsertShortUrl(shortInfo)
	if err != nil {
		return ""
	}

	return viper.GetString("short.prefix") + "/" + shortInfo.RedirectUrl
}

// FindShortByEntity 根据实体查询短连接
func FindShortByEntity(param entity.ShortURL) entity.ShortURL {
	return mapper.SelectShortUrlInfoByParam(param.ShortParam, param.BizType)
}
