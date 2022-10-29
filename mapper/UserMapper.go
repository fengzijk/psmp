package mapper

import (
	"go-psmp/config"
	"go-psmp/pojo/entity"
	"go-psmp/utils/json"
)

type UserMapper struct {
}

func (userMapper *UserMapper) FindUserByUsername(username string) entity.UserInfoEntity {
	var res entity.UserInfoEntity
	find := db.Model(&entity.UserInfoEntity{}).Where("username = ?", username).Find(&res)

	err := find.Error
	if err != nil {
		return res
	}

	config.Log.Info(json.ToJson(err))

	return res
}
