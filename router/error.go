package router

import "github.com/kataras/iris"

func RunError(app *iris.Application)  {
	//错误处理
	app.OnErrorCode(iris.StatusNotFound, NotFound)
	app.OnErrorCode(iris.StatusInternalServerError, InternalServerError)
}

func NotFound(ctx iris.Context) {
	ctx.View("404.html")
}

func InternalServerError(ctx iris.Context) {
	ctx.WriteString("凉凉,服务器错误")
}

