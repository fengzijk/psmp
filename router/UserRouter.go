package router

import (
	"github.com/gin-gonic/gin"
	"go-psmp/controller"
)

func UserLoginRouter(router *gin.Engine) {
	login := router.Group("/login")
	{
		login.POST("", controller.Login)
	}
}
