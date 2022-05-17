package main

import (
	"douyin_mine/config"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Run(iris.Addr(":" + config.AppConfig.GetString("server.port"))) //监听视频端口
}

func init() {

}
