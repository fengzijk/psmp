package service

import (
	"github.com/catinello/base62"
	"github.com/gin-gonic/gin"
	"math/rand"
	"short-url/mapper"
	"short-url/pojo/entity"
	"strings"
	"time"
)

type Response gin.H

const (
	path = "st"
)

// CreateShort 生成短连接
func CreateShort(longUrl string) string {

	var short entity.ShortURL
	short.LongUrl = longUrl

	urlInfoByEntity := mapper.SelectShortUrlInfoByEntity(short)

	if urlInfoByEntity.ShortUrl != "" {
		return short.ShortUrl
	}
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	sb.WriteString("/")
	sb.WriteString(path)
	sb.WriteString("/")
	timestamp := time.Now().UnixNano() / 1e6
	sb.WriteString(base62.Encode(int(timestamp)))
	shortUrl := sb.String()
	shortInfo := entity.ShortURL{
		ShortUrl: shortUrl,
		LongUrl:  longUrl,
	}
	err := mapper.InsertShortUrl(shortInfo)
	if err != nil {
		return ""
	}
	return shortUrl
}

// FindShortByEntity 根据实体查询短连接
func FindShortByEntity(param entity.ShortURL) entity.ShortURL {
	return mapper.SelectShortUrlInfoByEntity(param)
}
