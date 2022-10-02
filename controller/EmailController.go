package controller

import (
	"github.com/gin-gonic/gin"
	"go-psmp/pojo/request"
	"go-psmp/pojo/response"
	"go-psmp/service"
	"net/http"
)

func SendEmail(c *gin.Context) {

	emailRequest := request.SendEmailRequest{}
	err := c.BindJSON(&emailRequest)
	if err != nil {
		c.JSON(http.StatusAlreadyReported, *response.Fail("失败"))
	}

	to := []string{"guozhifengvip@163.com"}
	emailRequest.EmailTo = to
	service.SaveMail(emailRequest)
	resp := response.Result{
		Code: 200,
		Msg:  "OK",
		Data: "SUCCESS",
	}
	c.JSON(http.StatusOK, &resp)
}
