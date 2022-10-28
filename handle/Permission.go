package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-psmp/pojo/response"
	"go-psmp/utils/clamis"
)

func UserAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse, _ := clamis.GetClaims(c)

		obj := c.Request.URL.Path

		act := c.Request.Method

		sub := int(waitUse.AuthorityId)

		if sub == 0 {
			fmt.Print(obj, act)
			c.Next()
		} else {
			response.FailMessage(response.AuthorizationError, response.Text(response.AuthorizationError), c)
			c.Abort()
			return
		}
	}
}
