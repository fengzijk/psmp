package model

type EmailConfModel struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	ToUser   string `json:"toUser"`
	CcUser   string `json:"ccUser"`
}
