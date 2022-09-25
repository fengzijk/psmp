package controller

import (
	"github.com/gin-gonic/gin"
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

	shortParam := c.Param("pname")

	shortEntry := service.FindShortByByShortParam(shortParam)

	if shortEntry.RedirectUrl != "" && shortEntry.BizType == "url" {
		c.Redirect(http.StatusMovedPermanently, shortEntry.LongParam)
	}
}
