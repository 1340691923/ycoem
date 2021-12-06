package controller

import (
	"strings"

	"ycoem/lib/db"
	"ycoem/lib/logs"
	model "ycoem/lib/models"
	"ycoem/lib/request"
	"ycoem/lib/response"
	"ycoem/lib/service/abstract"
	"ycoem/lib/service/firm"
	"ycoem/lib/service/lunbo"
	"ycoem/lib/util"

	"github.com/kataras/iris"
	"go.uber.org/zap"
)

type FirmController struct {
	BaseController
}

func (this FirmController) IndexAction(ctx iris.Context) {
	const Tag = "FirmController"
	name := ctx.Params().Get("name")
	firmService := firm.FirmService{}
	options, err := abstract.Options("name", firmService)
	options = append(options, "公司介绍")
	if err != nil {
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		ctx.NotFound()
		return
	}

	if name == "" && len(options) > 0 {
		name = options[0]
	}

	if name == "公司介绍" {
		introduce, err := firmService.Get()
		if err != nil {
			logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
			ctx.NotFound()
			return
		}
		ctx.ViewData("introduce", introduce)
	} else {
		var firmModel []model.FirmModel
		firmService.Name = name
		err = abstract.SearchAll(firmService, &firmModel)
		if err != nil {
			logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
			ctx.NotFound()
			return
		}
		var firmList []response.FirmModel

		for _, v := range firmModel {

			tmpFrim := response.FirmModel{
				ID:         v.ID,
				Title:      v.Title,
				Img:        strings.Split(v.Img, ","),
				WxLink:     v.WxLink,
				Content:    v.Content,
				Name:       v.Name,
				CreateTime: v.CreateTime,
			}

			firmList = append(firmList, tmpFrim)
		}
		ctx.ViewData("firm", firmList)
	}

	lunboService := lunbo.LunboService{Moudle: "firm"}

	ctx.ViewData("options", options)
	ctx.ViewData("name", name)
	ctx.ViewData("lunboList", lunboService.GetLunboList())
	ctx.View("firm.html")
}

func (this FirmController) ListAction(ctx iris.Context) {
	reqData := request.Firm{}
	err := ctx.ReadJSON(&reqData)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	firmService := firm.FirmService{}
	firmService.Name = reqData.Name
	list := []model.FirmModel{}
	err = abstract.SearchList(firmService, reqData.Page, reqData.Limit, &list)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	count, err := abstract.Count(firmService)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, map[string]interface{}{"list": list, "count": count})
	return
}

func (this FirmController) NameOptionsAction(ctx iris.Context) {
	service := firm.FirmService{}
	list, err := abstract.Options("name", service)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, map[string]interface{}{"list": list})
	return
}

func (this FirmController) UpdateAction(ctx iris.Context) {
	reqData := request.Firm{}
	err := ctx.ReadJSON(&reqData)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	service := firm.FirmService{}
	service.Name = reqData.Name

	service.Img = strings.Join(reqData.Img, ",")

	service.Title = reqData.Title
	service.Content = reqData.Content
	service.WxLink = reqData.WxLink

	if reqData.ID == 0 {
		_, err = abstract.Insert(service)
	} else {
		err = abstract.Update(reqData.ID, service)
	}

	if err != nil {
		this.Error(ctx, err)
		return
	}
	db.Set(model.FirmCacheForIndex, "")
	this.Success(ctx, response.OperateSuccess, nil)
	return
}

func (this FirmController) DeleteAction(ctx iris.Context) {
	reqData := request.Firm{}
	err := ctx.ReadJSON(&reqData)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	serviceService := firm.FirmService{}
	err = abstract.Delete(reqData.ID, serviceService)

	if err != nil {
		this.Error(ctx, err)
		return
	}
	db.Set(model.FirmCacheForIndex, "")
	this.Success(ctx, response.OperateSuccess, nil)
	return
}

func (this FirmController) GetIntroduceAction(ctx iris.Context) {
	seo := firm.FirmService{}
	seoModel, err := seo.Get()
	if util.FilterRedisNilErr(err) {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.SearchSuccess, seoModel)
}

func (this FirmController) SetIntroduceAction(ctx iris.Context) {
	var model model.IntroduceModel

	err := ctx.ReadJSON(&model)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	seo := firm.FirmService{}
	err = seo.Save(model)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, "设置成功", nil)
}
