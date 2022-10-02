package router

import (
	"github.com/gin-gonic/gin"
	"go-psmp/handle"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.Use(handle.Cors(), gin.Recovery(), handle.Log())

	// 短连接路由
	ShortRouter(router)
	// 登录路由
	UserLoginRouter(router)
	// Email路由
	EmailRouter(router)
	// 首页路由
	HomePageRouter(router)

	// Agent路由
	AgentRouter(router)
	return router
}
