package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-psmp/pojo/request"
	"go-psmp/pojo/response"
	"go-psmp/service"
	"net/http"
)

const (
	CPU  = "cpu"
	DISK = "disk"
	APP  = "app"
)

var wxPushService = service.ServiceGroup.WxPushService

func SendEmail(c *gin.Context) {

	bizType := c.Param("bizType")
	dto := request.AlarmRequest{}

	err := c.BindJSON(&dto)
	if err != nil {
		response.Fail("失败", c)
	}

	notifyEmail(bizType, dto.Body)

	notifyWxPush(bizType, dto.Body)

	resp := response.ResponseResult{
		Code: 200,
		Msg:  "OK",
		Data: "SUCCESS",
	}
	c.JSON(http.StatusOK, &resp)

}

func notifyEmail(bizType, body string) {

	var fromName, toUser string

	if CPU == bizType {
		fromName = "CPU告警 "
		toUser = viper.GetString("alarm-email.cpu")
	}

	if DISK == bizType {
		fromName = "磁盘告警 "
		toUser = viper.GetString("alarm-email.disk")
	}

	if APP == bizType {
		fromName = "AppService告警 "
		toUser = viper.GetString("alarm-email.app")
	}

	emailRequest := request.SendEmailRequest{
		FromName: fromName,
		ToUser:   toUser,
		Body:     body,
		CcUser:   toUser,
	}

	emailService.SaveMail(emailRequest)
}

func notifyWxPush(bizType, body string) {

	partId := viper.GetString("alarm-weixin.toPartyId")
	toUser := viper.GetString("alarm-weixin.toUser")
	agentId := viper.GetInt("alarm-weixin.agentId")
	var fromName string

	if CPU == bizType {
		fromName = "CPU告警 "
	}

	if DISK == bizType {
		fromName = "磁盘告警 "
	}

	if APP == bizType {
		fromName = "AppService告警 "
	}

	wxPushService.SaveWxPushMessage(toUser, partId, fromName+body, agentId)
}
