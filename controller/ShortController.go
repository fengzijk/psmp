package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"short-url/pojo/entity"
	"short-url/service"
)

type Response gin.H

func CreateShort(c *gin.Context) {

	content := c.Param("param")
	fmt.Print(content)

	shortUrl := service.CreateShort(content)

	c.JSON(http.StatusOK, Response{
		"code": "200",
		"msg":  "成功",
		"data": Response{
			"short": viper.GetString("short.prefix") + shortUrl,
		},
	})
}

// 重定向
func Redirect() gin.HandlerFunc {
	return func(context *gin.Context) {
		url := context.Request.URL
		var short entity.ShortURL
		short.ShortUrl = url.String()
		result := service.FindShortByEntity(short)

		if result.LongUrl != "" {
			context.Redirect(http.StatusMovedPermanently, result.LongUrl)
		}
	}

}
