package router

import (
	"ycoem/middleware"

	"github.com/kataras/iris"
)

//go-bindata -debug -o=data/data.go -pkg=data view/... back/...       打包静态资源文件
//go-bindata  -o=data/data.go -pkg=data view/... back/...
func Run(app *iris.Application) {
	app.Use(middleware.Cors)
	RunTemplate(app)
	RunError(app)
	RunStatic(app)
	RunHome(app)
	RunAdmin(app)
}
