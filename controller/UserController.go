package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-psmp/pojo/request"
	"go-psmp/pojo/response"
	"go-psmp/service"
	"net/http"
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
