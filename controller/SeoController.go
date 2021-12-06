package controller

import (
	"github.com/kataras/iris"
	"ycoem/lib/request"
	"ycoem/lib/response"
	seo2 "ycoem/lib/service/seo"
	"ycoem/lib/util"
)

type SeoController struct {
	BaseController
}

func(this SeoController)GetAction(ctx iris.Context){
	seo := seo2.SeoService{}
	seoModel,err := seo.Get()
	if util.FilterRedisNilErr(err){
		this.Error(ctx,err)
		return
	}

	this.Success(ctx,response.SearchSuccess,seoModel)
}

func(this SeoController)CreateAction(ctx iris.Context){
	var seo request.Seo
	err := ctx.ReadJSON(&seo)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	seoService := seo2.SeoService{}
	err = seoService.Save(seo)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,"设置成功",nil)
}



func(this SeoController)GetPhoneQQAction(ctx iris.Context){
	seo := seo2.SeoService{}
	seoModel,err := seo.GetPhoneQQ()
	if util.FilterRedisNilErr(err){
		this.Error(ctx,err)
		return
	}

	this.Success(ctx,response.SearchSuccess,seoModel)
}

func(this SeoController)SetPhoneQQAction(ctx iris.Context){
	var model request.PhoneQQ
	err := ctx.ReadJSON(&model)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	seoService := seo2.SeoService{}
	err = seoService.SetPhoneQQ(model)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,"设置成功",nil)
}

func(this SeoController)GetWebinfoAction(ctx iris.Context){
	seo := seo2.SeoService{}
	seoModel,err := seo.GetWebinfo()
	if util.FilterRedisNilErr(err){
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.SearchSuccess,seoModel)
}

func(this SeoController)SetWebinfoAction(ctx iris.Context){
	var model request.Webinfo
	err := ctx.ReadJSON(&model)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	seoService := seo2.SeoService{}
	err = seoService.SetWebinfo(model)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,"设置成功",nil)
}

