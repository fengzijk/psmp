package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Log() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("请求时间   : %s \n请求IP     : %s  \n请求方式   : %s\n请求URL    :%s  \n状态码     : %d \n耗时       : %s \nUser-Agent :\"%s\" %s \n",
			param.TimeStamp.Format("2006-01-02 03:04:05"),
			param.ClientIP,
			param.Method,
			param.Path,
			//param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
