package controller

import (
	"errors"
	"time"

	"ycoem/lib/db"
	"ycoem/lib/jwt"
	"ycoem/lib/logs"
	model "ycoem/lib/models"
	"ycoem/lib/my_error"
	"ycoem/lib/response"
	"ycoem/lib/service/gm_user"
	"ycoem/lib/util"

	"github.com/kataras/iris"
)

type UserController struct {
	BaseController
}

func (this UserController) CleanUser(ctx iris.Context) {
	id := util.CtxFormIntDefault(ctx, "id", 0)
	if id == 0 {
		err = errors.New("用户id不能为空")
		this.Error(ctx, err)
		return
	}
	centUser := model.CentUserModel{Id: int64(id)}
	err = centUser.CleanUser()
	if err != nil {
		util.WriteErr(ctx, util.ErrUnknown)
		return
	}
	util.WriteErr(ctx, util.ErrNotError)
}

func (this UserController) Login(ctx iris.Context) {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	var gmUserService gm_user.GmUserService
	token, err := gmUserService.CheckLogin(username, password)
	if err != nil {
		logs.Logger.Sugar().Errorf("登陆失败", err)
		err = my_error.NewBusiness(gm_user.AUTH_ERROR, gm_user.ERROR_AUTH)
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, "登陆成功", map[string]interface{}{"token": token})
}

func (this UserController) UserInfo(ctx iris.Context) {
	var gmUserService gm_user.GmUserService

	token := this.GetToken(ctx)
	claims, err := jwt.ParseToken(token)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	info, err := gmUserService.GetRoleInfo(claims.RoleId)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, "登陆成功", map[string]interface{}{"roles": []string{"admin"}, "introduction": info.Description, "name": info.RoleName, "list": info.RoleList, "avatar": "http://uc-dzjs.lynlzqy.com/dzjs/view/touxiang.jpeg"})
}

func (this UserController) LogoutAction(ctx iris.Context) {
	conn := db.Redis.Get()
	defer conn.Close()
	token := this.GetToken(ctx)
	var claims *jwt.Claims
	claims, err = jwt.ParseToken(token)
	if err != nil {
		logs.Logger.Sugar().Errorf("token err", err)
		this.Error(ctx, errors.New("注销失败"))
		return
	}
	exptime := claims.ExpiresAt
	diffTime := exptime - time.Now().Unix()

	b, err := model.AddBanToken(token, diffTime)
	if b == false {
		logs.Logger.Sugar().Errorf("注销失败", b, err)
		this.Error(ctx, errors.New("注销失败"))
		return
	}
	if err != nil {
		logs.Logger.Sugar().Errorf("err", err)
		this.Error(ctx, err)
		return
	}
	model.IsBanToken(token)

	this.Success(ctx, response.LogoutSuccess, nil)
}

func (this UserController) UserListAction(ctx iris.Context) {
	var userModel model.GmUserModel
	list, err := userModel.Select()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, list)
}

func (this UserController) DeleteUserAction(ctx iris.Context) {
	var userModel model.GmUserModel
	userModel.ID = int32(this.FormIntDefault(ctx, "id", 0))
	err := userModel.Delete()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.DeleteSuccess, nil)
}

func (this UserController) GetUserByIdAction(ctx iris.Context) {
	var userModel model.GmUserModel
	var id = int32(this.FormIntDefault(ctx, "id", 0))
	userModel.ID = id
	gmUser, err := userModel.GetUserById()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, gmUser)
}

func (this UserController) UserUpdateAction(ctx iris.Context) {
	var userModel model.GmUserModel
	var id = int32(this.FormIntDefault(ctx, "id", 0))
	userModel.ID = id
	userModel.Realname = ctx.FormValue("realname")
	userModel.RoleId = int32(this.FormIntDefault(ctx, "role_id", 0))
	userModel.Password = ctx.FormValue("password")
	userModel.Username = ctx.FormValue("username")

	err := userModel.Update()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.OperateSuccess, nil)
}

func (this UserController) UserAddAction(ctx iris.Context) {
	var userModel model.GmUserModel

	userModel.Realname = ctx.FormValue("realname")
	userModel.RoleId = int32(this.FormIntDefault(ctx, "role_id", 0))
	userModel.Password = ctx.FormValue("password")
	userModel.Username = ctx.FormValue("username")

	id, err := userModel.Insert()
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, id)
}
