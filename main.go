package main

import (
	. "douyin_mine/config"
	. "douyin_mine/controller"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := NewApp()
	app.HandleDir("/douyin/video", "./video")
	app.HandleDir("/douyin/picture", "./picture")
	app.Run(iris.Addr(":"+AppConfig.GetString("server.port")), iris.WithoutPathCorrectionRedirection, iris.WithCharset("UTF-8")) //监听视频端口
}

func NewApp() *iris.Application {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, notFound)
	mvc.Configure(app.Party("/douyin/user"), func(app *mvc.Application) {
		app.Handle(new(UserController))
	})
	mvc.Configure(app.Party("/douyin/favorite"), func(app *mvc.Application) {
		app.Handle(new(FavoriteController))
	})
	mvc.Configure(app.Party("/douyin/comment"), func(app *mvc.Application) {
		app.Handle(new(CommentController))
	})
	mvc.Configure(app.Party("/douyin/publish"), func(app *mvc.Application) {
		app.Handle(new(PublicController))
	})
	mvc.Configure(app.Party("/douyin/relation"), func(app *mvc.Application) {
		app.Handle(new(RelationController))
	})
	mvc.Configure(app.Party("/douyin/feed"), func(app *mvc.Application) {
		app.Handle(new(FeedController))
	})
	//视频资源路径
	//app.Get("douyin/video/{year:string}/{month:string}/{filename:string}", func(context iris.Context) {
	//	path := AppConfig.GetString("resources.video.path") + string(os.PathSeparator) + context.Params().GetString("year") + string(os.PathSeparator) + context.Params().GetString("month") + string(os.PathSeparator) + context.Params().GetString("filename")
	//	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "%v", err)
	//	}
	//	content, err := ioutil.ReadAll(file)
	//	context.ResponseWriter().Write(content)
	//})

	//图片资源路径 default douyin/picture/0/0/default.jpg
	//app.Get("douyin/picture/{year:string}/{month:string}/{filename:string}", func(context iris.Context) {
	//	path := AppConfig.GetString("resources.picture.path") + string(os.PathSeparator) + context.Params().GetString("year") + string(os.PathSeparator) + context.Params().GetString("month") + string(os.PathSeparator) + context.Params().GetString("filename")
	//	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "%v", err)
	//	}
	//	content, err := ioutil.ReadAll(file)
	//	context.ResponseWriter().Write(content)
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
