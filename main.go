package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"go-psmp/config"
	"go-psmp/mapper"
	"go-psmp/pojo/model"
	"go-psmp/router"
	"go-psmp/task"
	"log"
	"reflect"
)

func init() {

	// 第一步  初始化配置
	initConfig()
	//
	task.InitTask()

	// 第二步 初始化数据库连接
	mapper.InitGormDB()

	// 第三步 初始化redis连接
	config.InitRedisDb()

	// 第四步 初始化邮件
	config.InitEmail()
}

func ValidateJSONDateType(field reflect.Value) interface{} {
	if field.Type() == reflect.TypeOf(model.LocalTime{}) {
		timeStr := field.Interface().(model.LocalTime).String()
		// 0001-01-01 00:00:00 是 go 中 time.Time 类型的空值
		// 这里返回 Nil 则会被 validator 判定为空值，而无法通过 `binding:"required"` 规则
		if timeStr == "0001-01-01 00:00:00" {
			return nil
		}
		return timeStr
	}
	return nil
}

func main() {

	initRouter := router.InitRouter()
	initRouter.LoadHTMLFiles("templates/index.html", "templates/favicon.ico")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册 model.LocalTime 类型的自定义校验规则
		v.RegisterCustomTypeFunc(ValidateJSONDateType, model.LocalTime{})
	}
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
