package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	"short-url/config"
	"short-url/controller"
	"short-url/mapper"
)

func init() {

	//第一步 设置配置文件目录
	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("未找到application文件")
		} else {
			log.Println("读取application文件错误")
		}

		log.Fatal(err)
	}

	// 第二步 初始化数据库连接
	err, _ := mapper.InitGormDB()
	if err != nil {
		log.Fatal(err)
	}

	// 第三步 初始化redis连接
	config.InitRedisDb()
}

func main() {

	router := gin.Default()
	router.Use(controller.Redirect(), gin.Recovery())
	shortcut := router.Group("/create")
	{
		shortcut.POST("/short/:param", controller.CreateShort)

	}
	// 测试redis
	//fmt.Println(redis.Get("12222"))
	//redis.SetObj("111111", entity.ShortURL{ShortUrl: "11111", LongUrl: "11111111111"})
	//redis.Set("1111122222", "111133333")
	//user := entity.ShortURL{}
	//redis.GetObj("11111", &user)
	//fmt.Println(user.LongUrl)
	_ = router.Run(fmt.Sprintf(":%s", viper.GetString("server.port")))

}
