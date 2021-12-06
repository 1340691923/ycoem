package controller

import (
	"ycoem/lib/logs"
	model "ycoem/lib/models"
	"ycoem/lib/request"
	"ycoem/lib/response"
	"ycoem/lib/service/abstract"
	"ycoem/lib/service/lunbo"
	"ycoem/lib/service/solution"

	"github.com/kataras/iris"
	"go.uber.org/zap"
)

type SolutController struct {
	BaseController
}

func(this SolutController)IndexAction(ctx iris.Context){
	const Tag = "SolutController"
	name := ctx.Params().Get("name")
	service := solution.SolutionService{}

	options,err := abstract.Options("name",service)
	if err!=nil{
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		ctx.NotFound()
		return
	}
	if name == "" && len(options) >0{
		name = options[0]
	}
	service.Name = name
	list := []model.SolutModel{}
	err = abstract.SearchAll(service,&list)
	if err!=nil {
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		ctx.NotFound()
		return
	}

	var solut response.Solution

	fatdService := solution.FatdService{}
	fatdmodel := []model.FatdModel{}

	for index := range list{
		fatdService.Fatd.SoluID = int(list[index].ID)
		abstract.SearchAll(fatdService,&fatdmodel)
		solut.Solut = list[index]
		solut.Fatd = fatdmodel
		fatdmodel = nil
	}
	lunboService := lunbo.LunboService{Moudle:service.TableName()}
	ctx.ViewData("lunboList",lunboService.GetLunboList())
	ctx.ViewData("solutionList",solut)
	ctx.ViewData("options",options)
	ctx.ViewData("name",name)
	ctx.View("solut.html")
}

func(this SolutController) ListAction(ctx iris.Context){
	reqData := request.Solut{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	service := solution.SolutionService{}
	service.Name = reqData.Name
	list := []model.SolutModel{}
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

	var solut response.Solution
	solutionList := []response.Solution{}
	fatdService := solution.FatdService{}
	fatdmodel := []model.FatdModel{}

	for index := range list{
		fatdService.Fatd.SoluID = int(list[index].ID)
		abstract.SearchAll(fatdService,&fatdmodel)
		solut.Solut = list[index]
		solut.Fatd = fatdmodel
		solutionList = append(solutionList, solut)
		fatdmodel = nil
	}

	this.Success(ctx,response.SearchSuccess,map[string]interface{}{"list":solutionList,"count":count})
	return
}

func(this SolutController) NameOptionsAction(ctx iris.Context){
	service := solution.SolutionService{}
	list,err:= abstract.Options("name",service)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.SearchSuccess,map[string]interface{}{"list":list})
	return
}

func(this SolutController)UpdateAction(ctx iris.Context){
	reqData := request.Solut{}
	err := ctx.ReadJSON(&reqData)

	if err!=nil{
		this.Error(ctx,err)
		return
	}
	service := solution.SolutionService{}
	service.SolutModel = reqData.Solut

	if reqData.Solut.ID == 0{
		var solutId int
		solutId,err = abstract.Insert(service)
		for index := range reqData.Fatd{
			reqData.Fatd[index].SoluID = solutId
		}
	}else{
		err = abstract.Update(int(reqData.Solut.ID),service)
	}

	if err!=nil{
		this.Error(ctx,err)
		return
	}

	fatdService := solution.FatdService{}

	abstract.DeleteByColumn("solu_id",reqData.Solut.ID,fatdService) // 删除掉 子表内容 重新写入

	for index:= range reqData.Fatd {
		fatdService.Fatd = reqData.Fatd[index]
		_,err = abstract.Insert(fatdService)
	}

	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.OperateSuccess,nil)
	return
}

func(this SolutController)DeleteAction(ctx iris.Context){
	reqData := request.Solut{}
	err := ctx.ReadJSON(&reqData)
	if err!=nil{
		this.Error(ctx,err)
		return
	}
	service := solution.SolutionService{}
	fatdService := solution.FatdService{}

	err = abstract.DeleteJoin(reqData.ID,service,fatdService,"solu_id")

	if err!=nil{
		this.Error(ctx,err)
		return
	}
	this.Success(ctx,response.OperateSuccess,nil)
	return
}
