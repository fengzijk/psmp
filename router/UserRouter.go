package router

import (
	"github.com/gin-gonic/gin"
	"go-psmp/controller"
)

func UserLoginRouter(router *gin.RouterGroup) {

	{
		router.POST("/login", controller.Login)
	}

}
