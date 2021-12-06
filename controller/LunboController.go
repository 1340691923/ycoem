package controller

import (
	"ycoem/lib/request"
	"ycoem/lib/response"
	"ycoem/lib/service/lunbo"

	"github.com/kataras/iris"
)

type LunboController struct {
	BaseController
}

func (this LunboController)ListAction(ctx iris.Context){
	lunboService := lunbo.LunboService{}
	rlt,err := lunboService.All()
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	lunboTableList := []response.LunboTable{}
	LunboTable := response.LunboTable{}
	for k,v := range rlt{
		LunboTable.Name = k
		LunboTable.Img= v
		lunboTableList  = append(lunboTableList, LunboTable)
	}
	this.Success(ctx,response.SearchSuccess,map[string]interface{}{"list":lunboTableList})
	return
}

func (this LunboController)SetAction(ctx iris.Context){
	reqData := request.LunboTable{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	lunboService := lunbo.LunboService{}
	err = lunboService.Set(reqData.Name,reqData.Img)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.OperateSuccess,nil)
	return
}

func (this LunboController)DelAction(ctx iris.Context){
	reqData := request.LunboTable{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	lunboService := lunbo.LunboService{}
	err = lunboService.Del(reqData.Name)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.OperateSuccess,nil)
	return
}

func(this LunboController) NameOptionsAction(ctx iris.Context){
	list:= lunbo.ModuleList
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.SearchSuccess,map[string]interface{}{"list":list})
	return
}

