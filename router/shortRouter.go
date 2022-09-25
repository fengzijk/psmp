package router

import (
	"github.com/gin-gonic/gin"
	"short-url/controller"
)

func ShortRouter(router *gin.Engine) {
	routerGroup := router.Group("/st")
	{
		routerGroup.POST("/create/:param", controller.CreateShort)
		router.GET("/:aa/:pname", controller.Redirect)
	}

}
