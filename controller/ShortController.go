package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"short-url/pojo/response"
	"short-url/service"
)

type Response gin.H

func CreateShort(c *gin.Context) {

	content := c.Param("param")
	shortUrl := service.CreateShort(content, "url")

	c.JSON(http.StatusOK, *response.Success(shortUrl))
}

// Redirect 重定向
func Redirect(c *gin.Context) {

	shortParam := c.Param("param")

	shortEntry := service.FindShortByByShortParam(shortParam)

	if shortEntry.RedirectUrl != "" && shortEntry.BizType == "url" {
		c.Redirect(http.StatusMovedPermanently, "https://"+shortEntry.LongParam)
	} else {
		c.Redirect(http.StatusMovedPermanently, viper.GetString("short.prefix"))
	}
}
