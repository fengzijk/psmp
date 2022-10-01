package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"short-url/enum"
	"short-url/pojo/request"
	"short-url/pojo/response"
	"short-url/service"
	_ "short-url/utils"
	"short-url/utils/monitor"
)

type Response gin.H

func CreateShort(c *gin.Context) {

	content := request.ShortContentRequest{}
	err := c.BindJSON(&content)
	if err != nil {
		c.JSON(http.StatusAlreadyReported, *response.Fail("失败"))
	}

	shortUrl := service.CreateShort(content.Content, enum.BizTypeEnum.GetMsg(enum.BizTypeEnum(content.BizType)))
	fmt.Print(monitor.GetCpuPercent())
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
			//c.Request.URL = shortEntry.LongParam
			//c.h
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
