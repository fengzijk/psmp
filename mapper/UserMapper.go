package mapper

import (
	"fmt"
	"go-psmp/pojo/entity"
)

func FindUserByUsername(username string) entity.User {
	var res entity.User
	find := db.Model(&entity.User{}).Where("username = ?", username).Find(&res)

	err := find.Error
	if err != nil {
		return res
	}

	fmt.Println(res)

	return res
}
