package router

import (
	"github.com/gin-gonic/gin"
	"go-psmp/controller"
)

func EmailRouter(router *gin.Engine) {
	routerGroup := router.Group("/email")
	{
		routerGroup.POST("/save", controller.SendEmail)
	}

}
