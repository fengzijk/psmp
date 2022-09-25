package request

type UserLoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
