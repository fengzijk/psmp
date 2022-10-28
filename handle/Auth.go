package handle

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go-psmp/config"
	"go-psmp/pojo/response"
	"go-psmp/service"
	"go-psmp/utils/date"
	psmpJwt "go-psmp/utils/jwt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var jwtService = service.ServiceGroup.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailMessageGin(response.CacheGetError, "未登录或非法访问", c)
			c.Abort()
			return
		}

		j := psmpJwt.NewJWT()
		// parseToken 解析token包含的信息
		claims, code := j.ParseToken(token)
		if code != response.SuccessCode {
			if code == response.TokenExpired {
				response.FailMessage(response.AuthorizationExpired, response.Text(response.AuthorizationExpired), c)
				c.Abort()
				return
			}
			response.FailMessage(response.TokenNotValidYet, response.Text(response.TokenNotValidYet), c)
			c.Abort()
			return
		}

		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := date.ParseDuration(config.JwtConf.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))

			RedisJwtToken := jwtService.GetRedisJWT(newClaims.Username)
			if len(RedisJwtToken) > 0 {
				fmt.Println(RedisJwtToken)
				//jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
			}

			// 无论如何都要记录当前的活跃状态
			_ = jwtService.SetRedisJWT(newToken, newClaims.Username)

		}
		c.Set("claims", claims)
		c.Next()
	}
}
