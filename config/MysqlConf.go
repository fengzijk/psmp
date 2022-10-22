package config

import (
	_ "database/sql"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitDb() (db *gorm.DB, err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", viper.Get("mysql.username"), viper.GetString("mysql.password"), viper.Get("mysql.host"), viper.GetString("mysql.port"), viper.GetString("mysql.database"))
	db, err = gorm.Open(mysql.Open(dsn),

		&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	//	err = db.AutoMigrate(&entity.UserInfoEntity{})
	if err != nil {
		return nil, err
	}
	return db, err
}
