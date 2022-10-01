package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"go-psmp/config"
	"go-psmp/mapper"
	"go-psmp/router"
	"go-psmp/task"
	"log"
)

func init() {

	// 第一步  初始化配置
	initConfig()

	task.InitTask()

	// 第二步 初始化数据库连接
	mapper.InitGormDB()

	// 第三步 初始化redis连接
	config.InitRedisDb()

	// 第四步 初始化邮件
	config.InitEmail()
}

func main() {

	initRouter := router.InitRouter()
	initRouter.LoadHTMLFiles("templates/index.html", "templates/favicon.ico")
	port := viper.GetString("server.port")
	fmt.Printf("监听端口:%s\n", port)
	_ = initRouter.Run(fmt.Sprintf(":%s", port))

}

func initConfig() {
	//第一步 设置配置文件目录
	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("未找到application文件")
		} else {
			log.Println("读取application文件错误")
		}

		log.Println(err)
	}

}
