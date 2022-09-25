package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"short-url/pojo/response"
	"short-url/service"
)

func Login(c *gin.Context) {

	//
	content := c.Param("param")

	fmt.Print(content)
	shortUrl := service.CreateShort(content)

	c.JSON(http.StatusOK, *response.Success(shortUrl))
}
