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

func SendEmail(c *gin.Context) {

	bizType := c.Param("bizType")
	emailRequest := request.SendEmailRequest{}

	err := c.BindJSON(&emailRequest)
	if err != nil {
		c.JSON(http.StatusAlreadyReported, *response.Fail("失败"))
	}

	if CPU == bizType {

		emailRequest.FromName = "CPU告警 "
		emailRequest.ToUser = viper.GetString("alarm-email.cpu")
	}

	if DISK == bizType {
		emailRequest.FromName = "磁盘告警 "
		emailRequest.ToUser = viper.GetString("alarm-email.disk")
	}

	if APP == bizType {
		emailRequest.FromName = "AppService告警 "
		emailRequest.ToUser = viper.GetString("alarm-email.app")
	}

	//if CPU==bizType {
	//	emailRequest.ToUser= viper.GetString("alarm-email.cpu")
	//}

	service.SaveMail(emailRequest)
	resp := response.Result{
		Code: 200,
		Msg:  "OK",
		Data: "SUCCESS",
	}
	c.JSON(http.StatusOK, &resp)
}
