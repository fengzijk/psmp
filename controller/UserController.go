package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"short-url/pojo/request"
	"short-url/pojo/response"
	"short-url/service"
)

func Login(c *gin.Context) {

	loginReq := request.UserLoginRequest{}
	err := c.BindJSON(&loginReq)
	if err != nil {
		c.JSON(http.StatusAlreadyReported, *response.Fail("失败"))
	}

	fmt.Print(loginReq)
	token := service.LoginByUsername(loginReq)

	c.JSON(http.StatusOK, *response.Success(token))
}
