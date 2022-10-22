package router

import (
	"github.com/gin-gonic/gin"
	"go-psmp/controller"
)

func EmailRouter(router *gin.RouterGroup) {
	routerGroup := router.Group("/email-records")
	{
		routerGroup.POST("/:bizType/save", controller.SendEmail)
		routerGroup.GET("/list-page", controller.ListPageEmailByAdmin)
	}
}
