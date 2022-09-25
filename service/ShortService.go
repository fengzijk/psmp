package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"short-url/mapper"
	"short-url/pojo/entity"
	shortUtil "short-url/utils/short"
)

type Response gin.H

const (
	path = "st"
)

// CreateShort 生成短连接
func CreateShort(param string, bizType string) string {

	md5Code := shortUtil.Get16MD5Encode(param)
	shortParam := shortUtil.GetShortParam(param)
	urlInfoByEntity := mapper.SelectShortUrlInfoByMd5Code(md5Code)

	if urlInfoByEntity.RedirectUrl != "" && urlInfoByEntity.BizType == "url" {
		return viper.GetString("short.prefix") + "/" + urlInfoByEntity.RedirectUrl
	}

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
