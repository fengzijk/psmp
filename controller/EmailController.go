package controller

import (
	"github.com/gin-gonic/gin"
	"go-psmp/pojo/response"
	"go-psmp/service"
	"strconv"
)

var emailService = service.ServiceGroup.EmailService

func ListPageEmailByAdmin(c *gin.Context) {

	status := c.Param("status")
	pageSize := c.Query("pageSize")
	pageNum := c.Query("pageNum")
	pageNumInt, _ := strconv.Atoi(pageNum)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	res := emailService.ListPageEmailByAdmin(status, pageNumInt, pageSizeInt)
	response.Ok(res, c)
}
