package model

type EmailConfModel struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
}
