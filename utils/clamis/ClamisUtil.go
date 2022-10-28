package clamis

import (
	"github.com/gin-gonic/gin"
	"go-psmp/pojo/request"
	"go-psmp/pojo/response"
	"go-psmp/utils/jwt"
)

func GetClaims(c *gin.Context) (*request.CustomClaims, int) {
	token := c.Request.Header.Get("x-token")
	j := jwt.NewJWT()
	claims, code := j.ParseToken(token)
	if code != 200 {

	}
	return claims, code
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != response.SuccessCode {
			return 0
		} else {
			return cl.UserId
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.UserId
	}
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != response.SuccessCode {
			return ""
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.UUID
	}
}

// GetUserAuthorityId 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserAuthorityId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != response.SuccessCode {
			return 0
		} else {
			return cl.AuthorityId
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.AuthorityId
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *request.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != response.SuccessCode {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse
	}
}
