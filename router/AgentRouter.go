package router

import (
	"github.com/gin-gonic/gin"
	"go-psmp/controller"
)

func AgentRouter(router *gin.Engine) {
	login := router.Group("/open/agent")
	{
		login.GET("/heartbeat/:agentIp/:agentName", controller.AgentHeartbeat)
	}
}
