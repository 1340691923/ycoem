package controller

import (
	"github.com/kataras/iris"
	model "ycoem/lib/models"
	"ycoem/lib/response"
	"ycoem/lib/service/gm_role"
)

type RoleController struct {
	BaseController
}

func (this RoleController) RolesAction(ctx iris.Context) {
	var service gm_role.GmRoleService
	list, err := service.Select()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, list)
}

func (this RoleController) RolesAddAction(ctx iris.Context) {

	var model2 model.GmRoleModel

	err = ctx.ReadJSON(&model2)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	var service gm_role.GmRoleService
	id, err := service.Add(model2)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, map[string]interface{}{"id": id})
}

func (this RoleController) RolesUpdateAction(ctx iris.Context) {
	var model2 model.GmRoleModel
	err = ctx.ReadJSON(&model2)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	var service gm_role.GmRoleService
	err = service.Update(model2)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
}

func (this RoleController) RolesDelAction(ctx iris.Context) {
	id := this.FormIntDefault(ctx, "id", 0)

	var service gm_role.GmRoleService
	err = service.Delete(id)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
}

func (this RoleController) RoleOptionAction(ctx iris.Context) {

	var model model.GmRoleModel

	list, err := model.Select()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, list)
}
