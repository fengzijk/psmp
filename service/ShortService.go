package service

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"short-url/pojo/entity"
)

type Response gin.H

const (
	path = "st"
)

// CreateShort 生成短连接
func CreateShort(longUrl string) string {

	//var short entity.ShortURL
	//short.LongUrl = longUrl
	//
	//urlInfoByEntity := mapper.SelectShortUrlInfoByMd5Code(short)
	//
	//if urlInfoByEntity.ShortUrl != "" {
	//	return viper.GetString("short.prefix") + short.ShortUrl
	//}
	//rand.Seed(time.Now().UnixNano())
	//var sb strings.Builder
	//sb.WriteString("/")
	//sb.WriteString(path)
	//sb.WriteString("/")
	//timestamp := time.Now().UnixNano() / 1e6
	//sb.WriteString(base62.Encode(int(timestamp)))
	//shortUrl := sb.String()
	//shortInfo := entity.ShortURL{
	//	ShortUrl: shortUrl,
	//	LongUrl:  longUrl,
	//}
	//err := mapper.InsertShortUrl(shortInfo)
	//if err != nil {
	//	return ""
	//}

	return viper.GetString("short.prefix") + longUrl
}

// FindShortByEntity 根据实体查询短连接
func FindShortByEntity(param entity.ShortURL) entity.ShortURL {
	//return mapper.SelectShortUrlInfoByEntity(param)
	return param
}
