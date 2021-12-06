package controller

import (
	"ycoem/lib/db"
	"ycoem/lib/logs"
	model "ycoem/lib/models"
	"ycoem/lib/service/abstract"
	"ycoem/lib/service/cases"
	"ycoem/lib/service/firm"
	"ycoem/lib/service/lunbo"
	"ycoem/lib/util"

	jsoniter "github.com/json-iterator/go"
	"github.com/kataras/iris"
	"go.uber.org/zap"
)

var err error

type IndexController struct {

}

//首页打开速度得尽可能的快 上缓存
func(this IndexController) IndexAction(ctx iris.Context){
	const Tag  =  "IndexAction"
	casesService := cases.CasesService{}
	caseslist := []model.CasesModel{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	casesCacheForIndex := db.Get(model.CasesCacheForIndex)
	//案例 列表
	if casesCacheForIndex==""{
		err = abstract.SearchAll(casesService,&caseslist)
		if err != nil {
			logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
			ctx.NotFound()
			return
		}
		go func() {
			str,_ := json.Marshal(caseslist)
			db.Set(model.CasesCacheForIndex,util.Bytes2str(str))
		}()
	}else{
		json.Unmarshal([]byte(casesCacheForIndex),&caseslist)
	}

	firmService := firm.FirmService{}
	firmList := []model.FirmModel{}
	firmCacheForIndex := db.Get(model.FirmCacheForIndex)
	//案例 列表
	if firmCacheForIndex==""{
		firmService.WxLink = "not null"
		err = abstract.SearchList(firmService,1,6,&firmList)
		if err != nil {
			logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
			ctx.NotFound()
			return
		}
		go func() {
			str,_ := json.Marshal(firmList)
			db.Set(model.FirmCacheForIndex,util.Bytes2str(str))
		}()
	}else{
		json.Unmarshal([]byte(firmCacheForIndex),&firmList)
	}

	ctx.ViewData("firmList",firmList)
	lunboService := lunbo.LunboService{Moudle:"index"}
	lunboList := lunboService.GetLunboList()
	ctx.ViewData("lunboList",lunboList)
	ctx.ViewData("cases",caseslist)
	ctx.View("index.html")
}



