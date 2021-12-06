package router

import (
	. "ycoem/controller"

	"github.com/kataras/iris"
)

//前台路由
func RunHome(app *iris.Application)  {
	//路由层
	app.Get("/", IndexController{}.IndexAction)          //首页	管理
	app.Get("/cases", CasesController{}.IndexAction)     //案例	管理
	app.Get("/cases/{name}", CasesController{}.IndexAction)     //案例	管理
	app.Get("/firm", FirmController{}.IndexAction)       //关于音创 管理
	app.Get("/firm/{name}", FirmController{}.IndexAction)       //关于音创 管理
	app.Get("/product/{name}/{nodeName}", ProductController{}.IndexAction) //产品	管理
	app.Get("/product/{name}", ProductController{}.IndexAction) //产品	管理
	app.Get("/service", ServiceController{}.IndexAction) //服务支持	管理
	app.Get("/service/{name}", ServiceController{}.IndexAction) //服务支持	管理
	app.Get("/solution", SolutController{}.IndexAction)     //解决方案	管理
	app.Get("/solution/{name}", SolutController{}.IndexAction)     //解决方案	管理
	app.Get("/concat", ConcatController{}.IndexAction) //联系我们	管理
}
