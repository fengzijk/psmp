package mapper

import (
	"gorm.io/gorm"
	"short-url/config"
)

var db *gorm.DB

func InitGormDB() (err error, err2 error) {
	db, _ = config.InitDb()
	return err, err2
}
