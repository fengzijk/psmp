package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomePageRouter(router *gin.Engine) {
	routerGroup := router.Group("/")
	{
		routerGroup.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Main website",
			})
		})

		routerGroup.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "",
			})
		})
	}

}
