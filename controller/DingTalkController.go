package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-psmp/pojo/response"
	"go-psmp/service"
	"net/http"
)

var dingTalkService = service.ServiceGroup.DingTalkService

func SenDingTalkMessage(c *gin.Context) {

	password := c.Query("password")

	message := c.Query("message")

	at := c.Query("at")
	resp := response.ResponseResult{
		Code: 200,
		Msg:  "OK",
		Data: "SUCCESS",
	}
	if password != viper.GetString("dingTalk.password") {
		c.JSON(http.StatusOK, &resp)
		return
	}

	dingTalkService.SendMessage(message, at)

	c.JSON(http.StatusOK, &resp)
}
