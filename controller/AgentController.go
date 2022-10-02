package controller

import (
	"github.com/gin-gonic/gin"
	"go-psmp/pojo/response"
	"go-psmp/service"
	"net/http"
)

func AgentHeartbeat(c *gin.Context) {

	agentIp := c.Param("agentIp")

	agentName := c.Param("agentName")
	//if strings.ContainsAny(agent,"-"){}

	service.Heartbeat(agentIp, agentName)
	resp := response.Result{
		Code: 200,
		Msg:  "OK",
		Data: "SUCCESS",
	}
	c.JSON(http.StatusOK, &resp)
}
