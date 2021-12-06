package controller

import (
	model "ycoem/lib/models"
	"ycoem/lib/request"
	"ycoem/lib/response"
	"ycoem/lib/service/abstract"
	"ycoem/lib/service/service"

	"github.com/kataras/iris"
)

type ServiceAnswerController struct {
	BaseController
}


func(this ServiceAnswerController) ListAction(ctx iris.Context){
	reqData := request.ServiceAnswer{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	serviceService := service.AnswerService{}
	serviceService.ParentID = reqData.ParentID
	list := []model.ServiceAnswerModel{}
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

func(this ServiceAnswerController) NameOptionsAction(ctx iris.Context){
	service := service.ServiceService{}
	service.Type = "解答"
	list,err:= abstract.Options("name",service)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.SearchSuccess,map[string]interface{}{"list":list})
	return
}

func(this ServiceAnswerController)UpdateAction(ctx iris.Context){
	reqData := request.ServiceAnswer{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	service := service.AnswerService{}
	service.Title = reqData.Title
	service.Img = reqData.Img
	service.Answer = reqData.Answer
	service.ParentID = reqData.ParentID

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

func(this ServiceAnswerController)DeleteAction(ctx iris.Context){
	reqData := request.ServiceAnswer{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}

	serviceService := service.AnswerService{}
	err = abstract.Delete(reqData.ID,serviceService)

	if err != nil {
		this.Error(ctx,err)
		return
	}

	this.Success(ctx,response.OperateSuccess,nil)
	return
}
