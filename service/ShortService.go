package service

import (
	"fmt"
	"go-psmp/enum"
	"go-psmp/mapper"
	"go-psmp/pojo/entity"
	"go-psmp/utils/redis"
	"go-psmp/utils/short"

	"github.com/spf13/viper"

	"strings"
)

const (
	path       = "st"
	cacheKey   = "short:md5code:"
	httpPrefix = "http://"
	https      = "https://"

	cacheShortKey = "short:key:"
)

type ShortService struct {
}

var shortMapper = mapper.MapperGroup.ShortUrlRecordMapper

// CreateShort 生成短连接
func (shortService *ShortService) CreateShort(param string, bizType string) string {

	var result string
	var urlEntity entity.ShortURLEntity

	// 参数校验
	if bizType == "" {
		bizType = enum.BizTypeEnum.GetMsg(1)
	}

	// https 截取
	if strings.HasPrefix(param, httpPrefix) || strings.HasPrefix(param, https) {
		bizType = enum.BizTypeEnum.GetMsg(2)
	}

	md5Code := short.Get16MD5Encode(param)
	shortParam := short.GetShortParam(param)

	redisCacheKey := cacheKey + md5Code

	// 缓存中查
	redis.GetObj(redisCacheKey, &urlEntity)

	// 数据库查询
	if urlEntity.LongParam == "" {
		urlEntity = shortMapper.SelectShortUrlInfoByMd5Code(md5Code)
	}

	// 存在返回结果
	if urlEntity.ShortParam != "" {
		redis.SetObj(redisCacheKey, urlEntity, 0)

		if urlEntity.BizType == enum.BizTypeEnum.GetMsg(2) {
			result = fmt.Sprintf("%s/%s/%s", viper.GetString("short.prefix"), path, urlEntity.ShortParam)

		} else {
			result = urlEntity.ShortParam
		}

		return result
	}

	// 新增
	shortInfo := entity.ShortURLEntity{
		ShortParam: shortParam,
		LongParam:  param,
		Md5Code:    md5Code,
		BizType:    bizType,
	}
	err := shortMapper.InsertShortUrl(shortInfo)
	if err != nil {
		return ""
	}

	if shortInfo.BizType == enum.BizTypeEnum.GetMsg(2) {
		// 链接
		result = fmt.Sprintf("%s/%s/%s", viper.GetString("short.prefix"), path, shortParam)

	} else {
		result = shortParam
	}

	return result
}

// FindShortByByShortParam 根据实体查询短连接
func (shortService *ShortService) FindShortByByShortParam(shortParam string) entity.ShortURLEntity {

	var urlEntity entity.ShortURLEntity

	redisCacheKey := cacheShortKey + shortParam
	redis.GetObj(redisCacheKey, &urlEntity)

	if urlEntity.ID != 0 {
		return urlEntity
	} else {
		urlEntity = shortMapper.SelectShortUrlInfoByShortParam(shortParam)
	}

	if urlEntity.ID != 0 {
		redis.SetObj(redisCacheKey, urlEntity, 0)
	}

	return urlEntity
}
