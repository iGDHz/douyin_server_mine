package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
	"time"
)

var AppConfig = viper.New() //配置文件信息
var Log *log.Logger

func init() {
	filename := strconv.Itoa(time.Now().Year()) + "-" + time.Now().Month().String() + ".log"
	logfile, err := os.Create(filename)
	defer logfile.Close()
	if err != nil {
		fmt.Println("创建log文件错误")
	}
	Log = log.New(logfile, "[Debug]", log.Llongfile)
	Log.SetPrefix("[Info]")
	Log.SetFlags(Log.Flags() | log.LstdFlags)
	Log.Println("开始加载 application.yaml 配置文件......")
	AppConfig.AddConfigPath(".")           //设置配置文件路径
	AppConfig.SetConfigName("application") //设置配置文件名
	AppConfig.SetConfigType("yaml")        //设置配置文件类型
	log.Println("配置文件加载完成")
	if err := AppConfig.ReadInConfig(); err != nil {
		panic(err) //读取配置文件失败
	}
}
