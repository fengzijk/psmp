package service

import (
	"go-psmp/mapper"
	"go-psmp/pojo/request"
)

func LoginByUsername(loginReq request.UserLoginRequest) string {
	res := ""
	if loginReq.UserName == "" || loginReq.Password == "" {
		return res
	}
	user := mapper.FindUserByUsername(loginReq.UserName)

	if user.ID == 0 || user.Name == "" {
		return res
	}

	return "111111111111111111111111111111"
}
