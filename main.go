package main

import (
	"douyin_mine/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := NewApp()
	app.Run(iris.Addr(":" + config.AppConfig.GetString("server.port"))) //监听视频端口
}

func NewApp() *iris.Application {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, notFound)
	mvc.Configure(app.Party("/douyin/user"), func(app *mvc.Application) {
		app.Handle(new(UserController))
	})
	//mvc.Configure(app.Party("/douyin/favorite"), func(app *mvc.Application) {
	//	app.Handle(new(FavoriteController))
	//})
	//mvc.Configure(app.Party("/douyin/comment"), func(app *mvc.Application) {
	//	app.Handle(new(CommentController))
	//})
	//mvc.Configure(app.Party("/douyin/publish"), func(app *mvc.Application) {
	//	app.Handle(new(PublishController))
	//})
	//mvc.Configure(app.Party("/douyin/relation"), func(app *mvc.Application) {
	//	app.Handle(new(RelationController))
	//})
	//mvc.Configure(app.Party("/feed"), func(app *mvc.Application) {
	//	app.Handle(new(FeedController))
	//})
	return app
}

func notFound(ctx iris.Context) {
	code := ctx.GetStatusCode()
	msg := "404 Not Found"
	ctx.JSON(iris.Map{
		"Message": msg,
		"Code":    code,
	})
}
