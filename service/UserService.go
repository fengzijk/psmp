package service

import (
	"go-psmp/config"
	"go-psmp/mapper"
	"go-psmp/pojo/entity"
	"go-psmp/pojo/request"
	"go-psmp/utils/short"
)

func LoginByUsername(loginReq request.UserLoginRequest) entity.UserInfoEntity {
	if loginReq.UserName == "" || loginReq.Password == "" {
		return entity.UserInfoEntity{}
	}
	user := mapper.FindUserByUsername(loginReq.UserName)

	if user.ID == 0 || user.Username == "" {
		return user
	}

	passwordStr := short.GetMd5Code(loginReq.Password + user.PasswordSalt)
	if user.Password != passwordStr {
		config.Log.DPanic("密码错误")
	}

	return user
}
