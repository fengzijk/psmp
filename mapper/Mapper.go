package mapper

import (
	"go-psmp/config"
	"gorm.io/gorm"
	"log"
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

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 20:
			pageSize = 20
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
