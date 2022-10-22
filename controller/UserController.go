package controller

import (
	"github.com/gin-gonic/gin"
	"go-psmp/config"
	"go-psmp/pojo/entity"
	"go-psmp/pojo/request"
	"go-psmp/pojo/response"
	"go-psmp/service"
	"go-psmp/utils/jwt"
	"go.uber.org/zap"
)

var jwtService = service.ServiceGroup.JwtService

func Login(c *gin.Context) {

	loginReq := request.UserLoginRequest{}
	err := c.BindJSON(&loginReq)
	if err != nil {
		response.Fail("失败", c)
	}

	user := service.LoginByUsername(loginReq)

	if user.ID == 0 {
		response.Fail("账号或密码错误", c)
		return
	}

	BuildToken(c, user)

}

// jwt
func BuildToken(c *gin.Context, user entity.UserInfoEntity) {
	j := &jwt.JWT{SigningKey: []byte(config.JwtConf.SigningKey)} // 唯一签名
	claims := j.CreateClaims(request.BaseClaims{
		UUID:        user.Username,
		NickName:    user.Nickname,
		Username:    user.Username,
		AuthorityId: uint(user.AuthorityId),
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		config.Log.Error("获取token失败!", zap.Error(err))
		response.Fail("获取token失败", c)
		return
	}

	jwtStr := jwtService.GetRedisJWT(user.Username)

	if len(jwtStr) > 0 {
		response.Ok(response.UserLoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, c)
	} else {

		redisJWT := jwtService.SetRedisJWT(token, user.Username)
		if !redisJWT {
			config.Log.Error("设置登录状态失败!", zap.Error(err))
			response.Fail("设置登录状态失败", c)
			return
		}
		response.Ok(response.UserLoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, c)

	}
}
