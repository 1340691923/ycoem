package controller

import (
	"strings"

	"ycoem/lib/logs"
	model "ycoem/lib/models"
	"ycoem/lib/request"
	"ycoem/lib/response"
	"ycoem/lib/service/abstract"
	"ycoem/lib/service/product"
	"ycoem/lib/util"

	"github.com/kataras/iris"
	"go.uber.org/zap"
)

type ProductController struct {
	BaseController
}

func (this ProductController) IndexAction(ctx iris.Context) {
	const Tag = "ProductController"
	name := ctx.Params().Get("name")
	nodeName := ctx.Params().Get("nodeName")
	productService := product.ProductService{}
	productService.Name = name
	var productModel model.ProductModel
	err = abstract.FindOne(productService, &productModel)

	if err != nil {
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		ctx.NotFound()
		return
	}
	options, err := abstract.Options("name", productService)
	productNodeService := product.ProductNodeService{}

	productNodeService.ProductID = productModel.ID

	nodeOptions, err := abstract.Options("name", productNodeService)
	if err != nil {
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		ctx.NotFound()
		return
	}
	if nodeName == "" && len(nodeOptions) > 0 {
		nodeName = nodeOptions[0]
	}

	productNodeService.Name = nodeName
	var obj model.ProductNodeModel

	err = abstract.FindOne(productNodeService, &obj)

	if err != nil {
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		ctx.NotFound()
		return
	}

	var imgList []string
	if obj.Imgs != "" {
		imgList = strings.Split(obj.Imgs, ",")
	}
	var lunboList []response.Lunbo
	var lunbo response.Lunbo
	lunbo.Img = productModel.Img
	lunbo.Url = "/" + productService.TableName() + "/" + name
	lunboList = append(lunboList, lunbo)
	ctx.ViewData("lunboList", lunboList)

	ctx.ViewData("imgList", imgList)
	ctx.ViewData("name", name)
	ctx.ViewData("nodeName", nodeName)
	ctx.ViewData("options", options)
	ctx.ViewData("nodeOptions", nodeOptions)

	ctx.View("product.html")
}

func (this ProductController) ListAction(ctx iris.Context) {
	reqData := request.Product{}
	err := ctx.ReadJSON(&reqData)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	service := product.ProductService{}
	service.Name = reqData.Name
	list := []model.ProductModel{}
	err = abstract.SearchList(service, reqData.Page, reqData.Limit, &list)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	count, err := abstract.Count(service)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, map[string]interface{}{"list": list, "count": count})
	return
}

func (this ProductController) NameOptionsAction(ctx iris.Context) {
	service := product.ProductService{}
	list, err := abstract.Options("name", service)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, map[string]interface{}{"list": list})
	return
}

func (this ProductController) UpdateAction(ctx iris.Context) {
	reqData := request.Product{}
	err := ctx.ReadJSON(&reqData)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	service := product.ProductService{}
	service.Name = reqData.Name
	service.Img = reqData.Img

	if reqData.ID == 0 {
		_, err = abstract.Insert(service)
	} else {
		err = abstract.Update(reqData.ID, service)
	}

	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
	return
}

func (this ProductController) DeleteAction(ctx iris.Context) {
	reqData := request.Product{}
	err := ctx.ReadJSON(&reqData)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	service := product.ProductService{}
	err = abstract.Delete(reqData.ID, service)

	if err != nil {
		this.Error(ctx, err)
		return
	}
	productNodeService := product.ProductNodeService{}

	err = abstract.DeleteByColumn("product_id", reqData.ID, productNodeService)

	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
	return
}

func (this ProductController) FindNodeById(ctx iris.Context) {
	reqData := request.ProductNode{}
	err := ctx.ReadJSON(&reqData)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	service := product.ProductNodeService{}
	service.Name = reqData.Name
	service.ProductID = reqData.ProductID
	var productNodeModel model.ProductNodeModel
	err = abstract.FindOne(service, &productNodeModel)
	if util.FilterMysqlNilErr(err) {
		this.Error(ctx, err)
		return
	}
	reqData.ProductID = productNodeModel.ProductID
	reqData.ID = productNodeModel.ID
	reqData.Name = productNodeModel.Name
	if productNodeModel.Imgs != "" {
		reqData.Imgs = strings.Split(productNodeModel.Imgs, ",")
	} else {
		reqData.Imgs = make([]string, 0)
	}
	this.Success(ctx, response.SearchSuccess, reqData)
	return
}

func (this ProductController) AddProductNode(ctx iris.Context) {
	reqData := request.ProductNode{}
	err := ctx.ReadJSON(&reqData)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	service := product.ProductNodeService{}
	service.Name = reqData.Name
	service.ProductID = reqData.ProductID
	service.ID = reqData.ID

	service.Imgs = strings.Join(reqData.Imgs, ",")

	if service.ID == 0 {
		_, err = abstract.Insert(service)
	} else {
		err = abstract.Update(service.ID, service)
	}

	if util.FilterMysqlNilErr(err) {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
	return
}
