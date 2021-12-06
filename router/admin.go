package router

import (
	. "ycoem/controller"
	"ycoem/middleware"

	"github.com/kataras/iris"
)

//后台路由
func RunAdmin(app *iris.Application) {
	app.Any("/api/gm_user/login", UserController{}.Login)
	app.Any("/api/upload/img", UploadController{}.ImageAction)
	app.Any("/api/upload/zip", UploadController{}.ZipAction)
	app.Use(middleware.JwtMiddleware)
	gmUser(app)
	seo(app)
	concat(app)
	cases(app)
	solut(app)
	product(app)
	lunbo(app)
	service(app)
	serviceAnswer(app)
	serviceDown(app)
	firm(app)
	video(app)
}

func video(app *iris.Application) {
	video := app.Party("/api")
	{
		video.Any("GetSceneFlagAction", BanController{}.GetSceneFlagAction)
		video.Any("searchBanScene", BanController{}.SearchBanSceneAction)
		video.Any("SetSceneFlagAction", BanController{}.SetSceneFlagAction)
		video.Any("setBanScene", BanController{}.SetBanSceneAction)
		video.Any("DelBanScene", BanController{}.DelBanSceneAction)
		video.Any("CleanUser", UserController{}.CleanUser)

	}
}

func concat(app *iris.Application) {
	concat := app.Party("/api/concat")
	{
		concat.Any("list", ConcatController{}.ListAction)
		concat.Any("delete", ConcatController{}.DeleteAction)
		concat.Any("UpdateAction", ConcatController{}.UpdateAction)
		concat.Any("NameOptionsAction", ConcatController{}.NameOptionsAction)
		concat.Any("SetLogisticeAction", ConcatController{}.SetLogisticeAction)
		concat.Any("GetLogisticeAction", ConcatController{}.GetLogisticeAction)
		concat.Any("SetOfficeAction", ConcatController{}.SetOfficeAction)
		concat.Any("GetOfficeAction", ConcatController{}.GetOfficeAction)
		concat.Any("GetContactAction", ConcatController{}.GetContactAction)
		concat.Any("SetContactAction", ConcatController{}.SetContactAction)
	}
}

func cases(app *iris.Application) {
	concat := app.Party("/api/cases")
	{
		concat.Any("list", CasesController{}.ListAction)
		concat.Any("delete", CasesController{}.DeleteAction)
		concat.Any("UpdateAction", CasesController{}.UpdateAction)
		concat.Any("NameOptionsAction", CasesController{}.NameOptionsAction)
	}
}

func product(app *iris.Application) {
	router := app.Party("/api/product")
	{
		router.Any("list", ProductController{}.ListAction)
		router.Any("delete", ProductController{}.DeleteAction)
		router.Any("UpdateAction", ProductController{}.UpdateAction)
		router.Any("NameOptionsAction", ProductController{}.NameOptionsAction)
		router.Any("FindNodeById", ProductController{}.FindNodeById)
		router.Any("AddProductNode", ProductController{}.AddProductNode)
	}
}

func service(app *iris.Application) {
	router := app.Party("/api/service")
	{
		router.Any("list", ServiceController{}.ListAction)
		router.Any("delete", ServiceController{}.DeleteAction)
		router.Any("UpdateAction", ServiceController{}.UpdateAction)
		router.Any("NameOptionsAction", ServiceController{}.NameOptionsAction)
	}
}

func serviceDown(app *iris.Application) {
	router := app.Party("/api/serviceDown")
	{
		router.Any("list", ServiceDownController{}.ListAction)
		router.Any("delete", ServiceDownController{}.DeleteAction)
		router.Any("UpdateAction", ServiceDownController{}.UpdateAction)
		router.Any("NameOptionsAction", ServiceDownController{}.NameOptionsAction)
	}
}

func firm(app *iris.Application) {
	router := app.Party("/api/firm")
	{
		router.Any("list", FirmController{}.ListAction)
		router.Any("delete", FirmController{}.DeleteAction)
		router.Any("UpdateAction", FirmController{}.UpdateAction)
		router.Any("NameOptionsAction", FirmController{}.NameOptionsAction)
		router.Any("SetIntroduceAction", FirmController{}.SetIntroduceAction)
		router.Any("GetIntroduceAction", FirmController{}.GetIntroduceAction)
	}
}

func serviceAnswer(app *iris.Application) {
	router := app.Party("/api/serviceAnswer")
	{
		router.Any("list", ServiceAnswerController{}.ListAction)
		router.Any("delete", ServiceAnswerController{}.DeleteAction)
		router.Any("UpdateAction", ServiceAnswerController{}.UpdateAction)
		router.Any("NameOptionsAction", ServiceAnswerController{}.NameOptionsAction)
	}
}

func lunbo(app *iris.Application) {
	router := app.Party("/api/lunbo")
	{
		router.Any("list", LunboController{}.ListAction)
		router.Any("UpdateAction", LunboController{}.SetAction)
		router.Any("NameOptionsAction", LunboController{}.NameOptionsAction)
		router.Any("delete", LunboController{}.DelAction)
	}
}

func seo(app *iris.Application) {
	seo := app.Party("/api/seo")
	{
		seo.Any("create", SeoController{}.CreateAction)
		seo.Any("get", SeoController{}.GetAction)
		seo.Any("setphoneQQ", SeoController{}.SetPhoneQQAction)
		seo.Any("getphoneQQ", SeoController{}.GetPhoneQQAction)
		seo.Any("SetWebinfoAction", SeoController{}.SetWebinfoAction)
		seo.Any("GetWebinfoAction", SeoController{}.GetWebinfoAction)
	}
}

func solut(app *iris.Application) {
	seo := app.Party("/api/solut")
	{
		seo.Any("list", SolutController{}.ListAction)
		seo.Any("delete", SolutController{}.DeleteAction)
		seo.Any("UpdateAction", SolutController{}.UpdateAction)
		seo.Any("NameOptionsAction", SolutController{}.NameOptionsAction)
	}
}

func gmUser(app *iris.Application) {
	gmUser := app.Party("/api/gm_user")
	{
		gmUser.Any("info", UserController{}.UserInfo)
		gmUser.Any("roles", RoleController{}.RolesAction)
		gmUser.Any("role/update", RoleController{}.RolesUpdateAction)
		gmUser.Any("role/add", RoleController{}.RolesAddAction)
		gmUser.Any("role/delete", RoleController{}.RolesDelAction)
		gmUser.Any("logout", UserController{}.LogoutAction)
		gmUser.Any("userlist", UserController{}.UserListAction)
		gmUser.Any("roleOption", RoleController{}.RoleOptionAction)
		gmUser.Any("getUserById", UserController{}.GetUserByIdAction)
		gmUser.Any("UpdateUser", UserController{}.UserUpdateAction)
		gmUser.Any("InsertUser", UserController{}.UserAddAction)
		gmUser.Any("DelUser", UserController{}.DeleteUserAction)
	}
}
