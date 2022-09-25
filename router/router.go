package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())

	// 短连接路由
	ShortRouter(router)
	// 登录路由
	UserLoginRouter(router)
	// 缓存路由
	//RedisRouter(router)

	return router
}
