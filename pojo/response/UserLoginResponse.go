package response

import "go-psmp/pojo/entity"

type UserLoginResponse struct {
	User      entity.UserInfoEntity `json:"user"`
	Token     string                `json:"token"`
	ExpiresAt int64                 `json:"expiresAt"`
}
