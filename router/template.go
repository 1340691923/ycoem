package router

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/view"
	"ycoem/data"
	"ycoem/lib/tmplateFun"
)

func RunTemplate(app *iris.Application) {
	app.RegisterView(NewTemplate())
}

func NewTemplate() (html *view.HTMLEngine) {
	html = iris.HTML("view/template", ".html")
	html.Binary(data.Asset, data.AssetNames)

	html.AddFunc("GetSeo", tmplateFun.GetSeoObj)
	html.AddFunc("GetPhoneObj", tmplateFun.GetPhoneObj)
	html.AddFunc("GetJsonValue", tmplateFun.GetJsonValue)
	html.AddFunc("GetNavicat", tmplateFun.GetNavicat)
	html.AddFunc("DeleteLine", tmplateFun.DeleteLine)
	html.AddFunc("NumberTransfer", tmplateFun.NumberTransfer)
	html.AddFunc("GetWebInfo", tmplateFun.GetWebInfo)
	html.AddFunc("GetContacts", tmplateFun.GetContacts)
	html.AddFunc("TimeFormat", tmplateFun.TimeFormat)
	html.Reload(true)
	return html
}
