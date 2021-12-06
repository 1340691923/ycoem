package controller

import (
	model "ycoem/lib/models"
	"ycoem/lib/request"
	"ycoem/lib/response"
	"ycoem/lib/service/abstract"
	"ycoem/lib/service/service"

	"github.com/kataras/iris"
)

type ServiceDownController struct {
	BaseController
}

func(this ServiceDownController) ListAction(ctx iris.Context){
	reqData := request.ServiceDown{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	serviceService := service.DownService{}
	serviceService.ParentID = reqData.ParentID
	list := []model.ServiceDownload{}
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

func(this ServiceDownController) NameOptionsAction(ctx iris.Context){
	service := service.ServiceService{}
	service.Type = "文件下载"
	list,err:= abstract.Options("name",service)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.SearchSuccess,map[string]interface{}{"list":list})
	return
}

func(this ServiceDownController)UpdateAction(ctx iris.Context){
	reqData := request.ServiceDown{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	service := service.DownService{}
	service.ParentID = reqData.ParentID
	service.Title = reqData.Title
	service.Version = reqData.Version
	service.DownloadURL = reqData.DownloadURL

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

func(this ServiceDownController)DeleteAction(ctx iris.Context){
	reqData := request.ServiceDown{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}

	serviceService := service.DownService{}
	err = abstract.Delete(reqData.ID,serviceService)

	if err != nil {
		this.Error(ctx,err)
		return
	}

	this.Success(ctx,response.OperateSuccess,nil)
	return
}

