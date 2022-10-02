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

func CreateShort(c *gin.Context) {

	content := request.ShortContentRequest{}
	err := c.BindJSON(&content)
	if err != nil {
		c.JSON(http.StatusAlreadyReported, *response.Fail("失败"))
	}

	shortUrl := service.CreateShort(content.Content, enum.BizTypeEnum.GetMsg(enum.BizTypeEnum(content.BizType)))

	//to := []string{"guozhifengvip@163.com"}
	//var dto = request.SendEmailRequest{EmailTo: to, Subject: "aaaaa", Content: "http://baidu.com", SystemName: "monitor"}
	//service.SaveMail("gzf", "guozhifengvip@163.com", "alarm", "http://baidu.com", "html")
	//service.SaveMail(dto)
	resp := response.Result{
		Code: 200,
		Msg:  "OK",
		Data: shortUrl,
	}
	c.JSON(http.StatusOK, &resp)
}

// Redirect 重定向
func Redirect(c *gin.Context) {

	shortParam := c.Param("param")

	shortEntry := service.FindShortByByShortParam(shortParam)

	if shortEntry.LongParam != "" {

		if shortEntry.BizType == enum.BizTypeEnum.GetMsg(2) {
			c.Redirect(http.StatusMovedPermanently, shortEntry.LongParam)

		} else {
			resp := response.Result{
				Code: 200,
				Msg:  "OK",
				Data: shortEntry.LongParam,
			}
			c.JSON(http.StatusOK, &resp)
		}
	}
}
