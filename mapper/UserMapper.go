package mapper

import (
	"fmt"
	"go-psmp/pojo/entity"
)

func FindUserByUsername(username string) entity.UserInfoEntity {
	var res entity.UserInfoEntity
	find := db.Model(&entity.UserInfoEntity{}).Where("username = ?", username).Find(&res)

	err := find.Error
	if err != nil {
		return res
	}

	fmt.Println(res)

	return res
}
