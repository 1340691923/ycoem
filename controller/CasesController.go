package controller

import (
	"ycoem/lib/db"
	"ycoem/lib/logs"
	model "ycoem/lib/models"
	"ycoem/lib/request"
	"ycoem/lib/response"
	"ycoem/lib/service/abstract"
	"ycoem/lib/service/cases"
	"ycoem/lib/service/lunbo"
	"ycoem/lib/util"

	"github.com/kataras/iris"
	"go.uber.org/zap"
)

type CasesController struct {
	BaseController
}

func(this CasesController)IndexAction(ctx iris.Context){
	const Tag = "CasesController"
	name := ctx.Params().Get("name")
	casesService :=	cases.CasesService{}

	options,err := abstract.Options("name",casesService)
	if err!=nil{
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		ctx.NotFound()
		return
	}

	if name == "" && len(options) >0{
		name = options[0]
	}

	casesService.Type = name

	page := util.CtxFormIntDefault(ctx,"page",1)
	list := []model.CasesModel{}
	err = abstract.SearchList(casesService,page,10,&list)

	if err != nil{
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		ctx.NotFound()
		return
	}
	count,err := abstract.Count(casesService)
	if err != nil{
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		ctx.NotFound()
		return
	}

	lunboService := lunbo.LunboService{Moudle:casesService.TableName()}
	ctx.ViewData("lunboList",lunboService.GetLunboList())
	ctx.ViewData("list",list)
	ctx.ViewData("name",name)
	ctx.ViewData("page",page)
	ctx.ViewData("count",util.CreatePageCount(count,10))
	ctx.ViewData("options",options)
	ctx.View("cases.html")
}

func(this CasesController) ListAction(ctx iris.Context){
	reqData := request.Cases{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	service := cases.CasesService{
		Title:    reqData.Title,
		Addr: reqData.Addr,
		Type:    reqData.Type,
	}
	list := []model.CasesModel{}
	err = abstract.SearchList(service,reqData.Page,reqData.Limit,&list)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	count,err:= abstract.Count(service)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.SearchSuccess,map[string]interface{}{"list":list,"count":count})
	return
}

func(this CasesController) NameOptionsAction(ctx iris.Context){
	service := cases.CasesService{}
	list,err:= abstract.Options("name",service)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.SearchSuccess,map[string]interface{}{"list":list})
	return
}

func(this CasesController)UpdateAction(ctx iris.Context){
	reqData := request.Cases{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	service := cases.CasesService{
		ID:         int32(reqData.Id),
		Title:      reqData.Title,
		Logo:       reqData.Logo,
		Desc:      	reqData.Desc,
		Type:       reqData.Type,
		Addr:       reqData.Addr,
	}
	if reqData.Id == 0{
		_,err = abstract.Insert(service)
	}else{
		err = abstract.Update(reqData.Id,service)
	}

	if err!=nil{
		this.Error(ctx,err)
		return
	}
	db.Set(model.CasesCacheForIndex,"")
	this.Success(ctx,response.OperateSuccess,nil)
	return
}

func(this CasesController)DeleteAction(ctx iris.Context){
	reqData := request.Cases{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	service := cases.CasesService{}
	err = abstract.Delete(reqData.Id,service)

	if err!=nil{
		this.Error(ctx,err)
		return
	}
	db.Set(model.CasesCacheForIndex,"")
	this.Success(ctx,response.OperateSuccess,nil)
	return
}
