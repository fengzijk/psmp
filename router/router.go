package router

import (
	"github.com/gin-gonic/gin"
	"short-url/handle"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(handle.Cors(), gin.Recovery())

	// 短连接路由
	ShortRouter(router)
	// 登录路由
	UserLoginRouter(router)
	// 缓存路由
	//RedisRouter(router)

	return router
}
