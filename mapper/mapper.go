package mapper

import (
	"gorm.io/gorm"
	"log"
	"short-url/config"
)

var db *gorm.DB

func InitGormDB() {

	initDb, err := config.InitDb()
	if err != nil {
		log.Println("初始化mySql失败:{}", err)
		return
	}
	db = initDb

}
