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
	// 缓存路由
	//RedisRouter(router)
	HomePageRouter(router)
	return router
}
