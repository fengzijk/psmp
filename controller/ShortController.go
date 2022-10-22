package controller

import (
	"github.com/gin-gonic/gin"
	"go-psmp/enum"
	"go-psmp/pojo/request"
	"go-psmp/pojo/response"
	"go-psmp/service"

	"net/http"
)

type Response gin.H

var shortService = service.ServiceGroup.ShortService

func CreateShort(c *gin.Context) {

	content := request.ShortContentRequest{}
	err := c.BindJSON(&content)
	if err != nil {
		response.Fail("失败", c)
	}

	shortUrl := shortService.CreateShort(content.Content, enum.BizTypeEnum.GetMsg(enum.BizTypeEnum(content.BizType)))

	//to := []string{"guozhifengvip@163.com"}
	//var dto = request.SendEmailRequest{EmailTo: to, Subject: "aaaaa", Content: "http://baidu.com", SystemName: "monitor"}
	//service.SaveMail("gzf", "guozhifengvip@163.com", "alarm", "http://baidu.com", "html")
	//service.SaveMail(dto)

	response.Ok(shortUrl, c)
}

// Redirect 重定向
func Redirect(c *gin.Context) {

	shortParam := c.Param("param")

	shortEntry := shortService.FindShortByByShortParam(shortParam)

	if shortEntry.LongParam != "" {

		if shortEntry.BizType == enum.BizTypeEnum.GetMsg(2) {
			c.Redirect(http.StatusMovedPermanently, shortEntry.LongParam)

		} else {

			response.Ok(shortEntry.LongParam, c)
		}
	}
}
