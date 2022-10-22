package router

import (
	"github.com/gin-gonic/gin"
	"go-psmp/handle"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	router.Use(handle.Cors(), gin.Recovery(), handle.Log())

	// 首页路由
	HomePageRouter(router)
	OpenGroup := router.Group("/open")
	PrivateGroup := router.Group("")
	PrivateGroup.Use(handle.JWTAuth())
	// 短连接路由
	ShortRouter(PrivateGroup)
	// 登录路由
	UserLoginRouter(OpenGroup)
	// Email路由
	EmailRouter(PrivateGroup)

	// Agent路由
	AgentRouter(router)
	return router
}
