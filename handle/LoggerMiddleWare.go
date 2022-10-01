package handle

import (
	"github.com/gin-gonic/gin"
	"log"
)

func LoggerMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 此处省略若干行。。。
		// 注意：在Next之前不读取body的内容
		c.Next()
		params, _ := c.Get("params")
		// 自定义打印，此处只是简单举例
		log.Printf("Request params:%v", params)
	}
}
