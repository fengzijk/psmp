package dto

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Password int64  `json:"password"`
}

func Login(c *gin.Context) {
	json := User{}
	c.BindJSON(&json)
	log.Printf("%v", &json)
	c.JSON(http.StatusOK, gin.H{
		"name":     json.Name,
		"password": json.Password,
	})
}
