package router

import (
	"github.com/gin-gonic/gin"
	"go-psmp/controller"
)

func SendMessageRouter(router *gin.RouterGroup) {

	{
		router.GET("/message/send-ding-talk-message", controller.SenDingTalkMessage)

	}
}
