package controller

import (
	"ycoem/lib/db"
	model "ycoem/lib/models"
	"ycoem/lib/service/abstract"
	"ycoem/lib/service/lunbo"

	"github.com/kataras/iris"

	"ycoem/lib/request"
	"ycoem/lib/response"
	"ycoem/lib/service/concat"
	"ycoem/lib/util"
)

type ConcatController struct {
	BaseController
}

func(this ConcatController) IndexAction(ctx iris.Context){
	serviceDepService := concat.ServiceDepService{}
	logisticeService := concat.LogisticeService{}
	branchOfficeService := concat.BranchOfficeService{}

	serviceDep,err := serviceDepService.All()

	if util.FilterRedisNilErr(err){
		ctx.Write([]byte(err.Error()))
		return
	}

	logistice,err:= logisticeService.Get()
	if util.FilterRedisNilErr(err){
		ctx.Write([]byte(err.Error()))
		return
	}
	branchOffice,err:= branchOfficeService.Get()
	if util.FilterRedisNilErr(err){
		ctx.Write([]byte(err.Error()))
		return
	}
	lunboService := lunbo.LunboService{Moudle:"concat"}
	ctx.ViewData("lunboList",lunboService.GetLunboList())
	ctx.ViewData("logistice",logistice)
	ctx.ViewData("branchOffice",branchOffice)
	ctx.ViewData("serviceDep",serviceDep)
	ctx.View("contact.html")
}

func(this ConcatController) ListAction(ctx iris.Context){
	reqData := request.Contact{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	service := concat.ServiceDepService{
		Area:    reqData.Area,
		Manager: reqData.Manager,
		Name:    reqData.Name,
	}
	list := []model.ServiceDepModel{}
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

func(this ConcatController) NameOptionsAction(ctx iris.Context){
	service := concat.ServiceDepService{}
	list,err:= abstract.Options("name",service)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.SearchSuccess,map[string]interface{}{"list":list})
	return
}

func(this ConcatController)UpdateAction(ctx iris.Context){
	reqData := request.Contact{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	service := concat.ServiceDepService{
		Area:    reqData.Area,
		Manager: reqData.Manager,
		Name:    reqData.Name,
		Support: reqData.Support,
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
	db.Set(model.ServiceDepRedisModel,"")

	this.Success(ctx,response.OperateSuccess,nil)
	return
}

func(this ConcatController)DeleteAction(ctx iris.Context){
	reqData := request.Contact{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	service := concat.ServiceDepService{}
	err = abstract.Delete(reqData.Id,service)

	if err!=nil{
		this.Error(ctx,err)
		return
	}
	db.Set(model.ServiceDepRedisModel,"")
	this.Success(ctx,response.OperateSuccess,nil)
	return
}

func(this ConcatController)SetLogisticeAction(ctx iris.Context){
	reqData := request.Logistics{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	service := concat.LogisticeService{}
	err = service.Save(reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.OperateSuccess,nil)
	return
}

func(this ConcatController)GetLogisticeAction(ctx iris.Context){
	service := concat.LogisticeService{}
	data,err := service.Get()
	if util.FilterRedisNilErr(err){
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.SearchSuccess,data)
	return
}

func(this ConcatController)SetOfficeAction(ctx iris.Context){
	reqData := request.BranchOffice{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	service := concat.BranchOfficeService{}
	err = service.Save(reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.OperateSuccess,nil)
	return
}

func(this ConcatController)GetOfficeAction(ctx iris.Context){
	service := concat.BranchOfficeService{}
	data,err := service.Get()
	if util.FilterRedisNilErr(err) {
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.SearchSuccess,data)
	return
}

func(this ConcatController)GetContactAction(ctx iris.Context){
	contactsService:= concat.ContactsService{}
	seoModel,err := contactsService.Get()
	if util.FilterRedisNilErr(err){
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.SearchSuccess,seoModel)
}

func(this ConcatController)SetContactAction(ctx iris.Context){
	var model request.Contacts
	err := ctx.ReadJSON(&model)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	contactsService := concat.ContactsService{}
	err = contactsService.Set(model)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,"设置成功",nil)
}

