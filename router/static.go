package router

import (
	"ycoem/data"

	"github.com/kataras/iris"
)

func RunStatic(app *iris.Application){
	// 静态资源
	app.StaticEmbedded("/static", "./view/static", data.Asset, data.AssetNames)
	app.StaticEmbedded("/back", "./back", data.Asset, data.AssetNames)
	app.StaticWeb("image","./file/img")
	app.StaticWeb("download","./file/download")
}
