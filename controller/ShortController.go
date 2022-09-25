package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"short-url/pojo/entity"
	"short-url/pojo/response"
	"short-url/service"
	shortUtil "short-url/utils/short"
)

type Response gin.H

func CreateShort(c *gin.Context) {

	content := c.Param("param")
	shortUrl := service.CreateShort(content, "url")

	c.JSON(http.StatusOK, *response.Success(shortUrl))
}

// Redirect 重定向
func Redirect() gin.HandlerFunc {
	return func(context *gin.Context) {
		url := context.Request.URL
		var short entity.ShortURL
		short.ShortParam = shortUtil.GetMd5Code(url.String())
		result := service.FindShortByEntity(short)

		if result.RedirectUrl != "" {
			context.Redirect(http.StatusMovedPermanently, result.RedirectUrl)
		}
	}
}
