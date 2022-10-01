package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"short-url/enum"
	"short-url/mapper"
	"short-url/pojo/entity"
	"short-url/utils/redis"
	shortUtil "short-url/utils/short"
	"strings"
)

type Response gin.H

const (
	path     = "st"
	cacheKey = "short:"
	http     = "http://"
	https    = "https://"
)

// CreateShort 生成短连接
func CreateShort(param string, bizType string) string {

	var result string
	var urlEntity entity.ShortURLEntity

	// 参数校验
	if bizType == "" {
		bizType = enum.BizTypeEnum.GetMsg(1)
	}

	// https 截取
	if strings.HasPrefix(param, http) || strings.HasPrefix(param, https) {
		strings.ReplaceAll(param, http, "")
		strings.ReplaceAll(param, https, "")
	}

	md5Code := shortUtil.Get16MD5Encode(param)
	shortParam := shortUtil.GetShortParam(param)

	redisCacheKey := cacheKey + shortParam

	// 缓存中查
	redis.GetObj(redisCacheKey, &urlEntity)

	// 数据库查询
	if urlEntity.LongParam == "" {
		urlEntity = mapper.SelectShortUrlInfoByMd5Code(md5Code)
	}

	// 存在返回结果
	if urlEntity.ShortParam != "" && urlEntity.BizType == enum.BizTypeEnum.GetMsg(2) {
		redis.SetObj(redisCacheKey, urlEntity)
		// 新增
		result = fmt.Sprintf("%s/%s/%s", viper.GetString("short.prefix"), path, shortParam)
		return result
	}

	// 新增
	shortInfo := entity.ShortURLEntity{
		ShortParam: shortParam,
		LongParam:  param,
		Md5Code:    md5Code,
		BizType:    bizType,
	}
	err := mapper.InsertShortUrl(shortInfo)
	if err != nil {
		return ""
	}

	result = fmt.Sprintf("%s/%s/%s", viper.GetString("short.prefix"), path, shortParam)
	return result
}

// FindShortByByShortParam 根据实体查询短连接
func FindShortByByShortParam(shortParam string) entity.ShortURLEntity {

	var urlEntity entity.ShortURLEntity

	redisCacheKey := cacheKey + shortParam
	redis.GetObj(redisCacheKey, &urlEntity)

	if urlEntity.LongParam == "" || urlEntity.BizType == enum.BizTypeEnum.GetMsg(2) {
		urlEntity = mapper.SelectShortUrlInfoByShortParam(shortParam)
	}

	return urlEntity
}
