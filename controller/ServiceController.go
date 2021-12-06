package controller

import (
	"ycoem/lib/logs"
	model "ycoem/lib/models"
	"ycoem/lib/request"
	"ycoem/lib/response"
	"ycoem/lib/service/abstract"
	"ycoem/lib/service/lunbo"
	"ycoem/lib/service/service"

	"github.com/kataras/iris"
	"go.uber.org/zap"
)

type ServiceController struct {
	BaseController
}

func(this ServiceController)IndexAction(ctx iris.Context){
	const Tag = "ServiceController"
	name := ctx.Params().Get("name")

	serviceService := service.ServiceService{}

	options,err := abstract.Options("name",serviceService)
	if err!=nil{
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		ctx.NotFound()
		return
	}
	if name == "" && len(options) >0{
		name = options[0]
	}

	downService := service.DownService{}
	downService.ParentID = name
	downList := []model.ServiceDownload{}
	abstract.SearchAll(downService,&downList)

	answerService := service.AnswerService{}
	answerService.ParentID = name
	answerList := []model.ServiceAnswerModel{}
	abstract.SearchAll(answerService,&answerList)

	lunboService := lunbo.LunboService{Moudle:"service"}


	ctx.ViewData("downList",downList)
	ctx.ViewData("answerList",answerList)
	ctx.ViewData("options",options)
	ctx.ViewData("lunboList",lunboService.GetLunboList())
	ctx.ViewData("name",name)
	ctx.View("service.html")
}

func(this ServiceController) ListAction(ctx iris.Context){
	reqData := request.Service{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	serviceService := service.ServiceService{}
	serviceService.Name = reqData.Name
	list := []model.ServiceModel{}
	err = abstract.SearchList(serviceService,reqData.Page,reqData.Limit,&list)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	count,err:= abstract.Count(serviceService)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.SearchSuccess,map[string]interface{}{"list":list,"count":count})
	return
}

func(this ServiceController) NameOptionsAction(ctx iris.Context){
	service := service.ServiceService{}
	list,err:= abstract.Options("name",service)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.SearchSuccess,map[string]interface{}{"list":list})
	return
}

func(this ServiceController)UpdateAction(ctx iris.Context){
	reqData := request.Service{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	service := service.ServiceService{}
	service.Name = reqData.Name
	service.Type = reqData.Type

	if reqData.ID == 0{
		_,err = abstract.Insert(service)
	}else{
		err = abstract.Update(reqData.ID,service)
	}

	if err != nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.OperateSuccess,nil)
	return
}

func(this ServiceController)DeleteAction(ctx iris.Context){
	reqData := request.Service{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}

	serviceService := service.ServiceService{}
	err = abstract.Delete(reqData.ID,serviceService)

	if err != nil {
		this.Error(ctx,err)
		return
	}

	answerService := service.AnswerService{}

	err = abstract.DeleteByColumn("parent_id",reqData.Name,answerService)

	if err != nil{
		this.Error(ctx,err)
		return
	}

	downService := service.DownService{}

	err = abstract.DeleteByColumn("parent_id",reqData.Name,downService)

	if err != nil {
		this.Error(ctx,err)
		return
	}

	this.Success(ctx,response.OperateSuccess,nil)
	return
}
