package router

import (
	"github.com/gin-gonic/gin"
	"short-url/controller"
)

func UserLoginRouter(router *gin.Engine) {
	login := router.Group("/login")
	{
		login.POST("", controller.Login)
	}
}
