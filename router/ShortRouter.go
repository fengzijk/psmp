package router

import (
	"github.com/gin-gonic/gin"
	"go-psmp/controller"
)

func ShortRouter(router *gin.RouterGroup) {
	routerGroup := router.Group("/st")
	{
		routerGroup.POST("/create", controller.CreateShort)
		routerGroup.GET("/:param", controller.Redirect)
	}

}
