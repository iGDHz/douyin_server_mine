package config

import (
	"github.com/spf13/viper"
)

var AppConfig = viper.New() //配置文件信息

func init() {
	AppConfig.AddConfigPath(".")           //设置配置文件路径
	AppConfig.SetConfigName("application") //设置配置文件名
	AppConfig.SetConfigType("yaml")        //设置配置文件类型

	if err := AppConfig.ReadInConfig(); err != nil {
		panic(err) //读取配置文件失败
	}
}
