package controller

import (
	"github.com/gin-gonic/gin"
	"go-psmp/pojo/response"
	"go-psmp/service"
	"net/http"
)

// var dingTalkService = service.ServiceGroup.DingTalkService

func AgentHeartbeat(c *gin.Context) {

	agentIp := c.Param("agentIp")

	agentName := c.Param("agentName")
	//if strings.ContainsAny(agent,"-"){}

	service.Heartbeat(agentIp, agentName)
	resp := response.ResponseResult{
		Code: 200,
		Msg:  "OK",
		Data: "SUCCESS",
	}
	c.JSON(http.StatusOK, &resp)
}
