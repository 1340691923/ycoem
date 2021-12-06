package controller

import (
	"strconv"

	"github.com/kataras/iris"
	"ycoem/lib/db"
	"ycoem/lib/logs"
	model "ycoem/lib/models"
	"ycoem/lib/request"
	"ycoem/lib/response"
	"ycoem/lib/service/scene"
)

type BanController struct {
	request.Request
	response.Response
}

func (this BanController) DelBanSceneAction(ctx iris.Context) {
	id := this.FormIntDefault(ctx, "id", 0)
	bs := model.BanSceneModel{Id: id}
	types, channo, sences, err := bs.GetInfoById()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	conn := db.Redis.Get()
	defer conn.Close()
	b, err := scene.DeleteBan(conn, types, channo, sences)
	if err != nil {
		logs.Logger.Sugar().Errorf("b", b, err)
		this.Error(ctx, err)
		return
	}
	err = bs.Delete(id)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, "删除成功", nil)
}

func (this BanController) SetBanSceneAction(ctx iris.Context) {
	sceneType := this.FormIntDefault(ctx, "type", 1)
	scenes := ctx.FormValue("scene")
	channo := this.FormIntDefault(ctx, "channo", 1)

	bs := model.BanSceneModel{
		Types:  strconv.Itoa(sceneType),
		Channo: strconv.Itoa(channo),
		Scene:  scenes,
	}
	err = bs.Insert(sceneType, channo, scenes)
	if err != nil {
		this.Error(ctx, err)
		return
	}
	conn := db.Redis.Get()
	defer conn.Close()
	b, err := scene.AddBan(conn, ctx.FormValue("type"), ctx.FormValue("channo"), scenes)
	if err != nil {
		logs.Logger.Sugar().Errorf("b", b, err)
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, "新增成功", nil)
	return
}

func (this BanController) SearchBanSceneAction(ctx iris.Context) {
	page := this.FormIntDefault(ctx, "page", 1)
	limit := this.FormIntDefault(ctx, "limit", 10)
	types := this.FormIntDefault(ctx, "type", 0)
	channo := this.FormIntDefault(ctx, "channo", 0)

	var bs model.BanSceneModel

	data, count, err := bs.Search(page, limit, types, channo)

	if err != nil {
		this.Error(ctx, err)
		return
	}
	if data == nil {
		data = []interface{}{}
	}
	this.Success(ctx, "查询成功",
		map[string]interface{}{
			"list":  data,
			"count": count,
			"limit": limit,
			"page":  page,
		})
	return
}

func (this BanController) GetSceneFlagAction(ctx iris.Context) {
	ss := model.CentSceneFlagModel{}
	flag := ss.GetFlag()
	this.Success(ctx, "查询成功", map[string]interface{}{"flag": flag})
	return
}

func (this BanController) SetSceneFlagAction(ctx iris.Context) {
	flag := this.FormIntDefault(ctx, "flag", 1)
	var flagBool = false
	if flag == 1 {
		flagBool = true
	}
	ss := model.CentSceneFlagModel{Flag: flagBool}
	err := ss.SetFlag()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, "修改成功", nil)
	return
}
