package router

import (
	"github.com/gin-gonic/gin"
	"go-psmp/controller"
)

func UserLoginRouter(router *gin.Engine) {
	PrivateGroup := router.Group("/manager")

	//PrivateGroup.Use(handle.JWTAuth())
	{
		PrivateGroup.POST("/login", controller.Login)
	}

	{
		PrivateGroup.GET("/email-records/list-page", controller.ListPageEmailByAdmin)
	}
}
