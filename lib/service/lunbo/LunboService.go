package lunbo

import (
	"encoding/json"

	"github.com/garyburd/redigo/redis"

	"ycoem/lib/db"
	"ycoem/lib/logs"
	model "ycoem/lib/models"
	"ycoem/lib/response"
	"ycoem/lib/service/abstract"
	"ycoem/lib/service/product"
	"ycoem/lib/util"

	"go.uber.org/zap"
)

type LunboService struct {
	Moudle string
	Name   string
}

var ModuleList = []string{"案例", "联系我们", "关于我们", "服务支持", "解决方案"}

var ModuleMap = map[string]string{
	"cases":    "案例",
	"concat":   "联系我们",
	"firm":     "关于我们",
	"service":  "服务支持",
	"solution": "解决方案",
}

func (this *LunboService) All() (rlt map[string]string, err error) {
	conn := db.Redis.Get()
	defer conn.Close()
	rlt, err = redis.StringMap(conn.Do("hgetall", model.LunboModel))
	return
}

func (this *LunboService) Set(module, img string) (err error) {
	conn := db.Redis.Get()
	defer conn.Close()

	_, err = redis.Bool(conn.Do("hset", model.LunboModel, module, img))
	return
}

func (this *LunboService) Del(module string) (err error) {
	conn := db.Redis.Get()
	defer conn.Close()
	_, err = redis.Bool(conn.Do("hdel", model.LunboModel, module))
	return
}

func (this *LunboService) GetImg() (img string, err error) {
	conn := db.Redis.Get()
	defer conn.Close()
	img, err = redis.String(conn.Do("hget", model.LunboModel, ModuleMap[this.Moudle]))
	return
}

func (this *LunboService) GetLunboList() (lunboList []response.Lunbo) {
	if this.Moudle == "index" {
		lunboCacheForIndex := db.Get(model.LunboCacheForIndex)
		if lunboCacheForIndex == "" {
			lunboList = this.lunbo()
			go func() {
				str, _ := json.Marshal(lunboList)
				db.Set(model.LunboCacheForIndex, util.Bytes2str(str))
			}()
		} else {
			json.Unmarshal([]byte(lunboCacheForIndex), &lunboList)
		}
		return lunboList
	}
	return this.oneImg()
}

//只展现一张图
func (this *LunboService) oneImg() (lunboList []response.Lunbo) {

	lunbo := response.Lunbo{}
	img, err := this.GetImg()
	if err != nil {
		return
	}

	lunbo.Img = img
	lunbo.Url = "/" + this.Moudle
	lunboList = append(lunboList, lunbo)

	return
}

//用于轮播
func (this *LunboService) lunbo() (lunboList []response.Lunbo) {
	const Tag = "LunboService"
	productList := []model.ProductModel{}
	productService := product.ProductService{}
	err := abstract.SearchAll(productService, &productList)
	if err != nil {
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		return
	}
	lunbo := response.Lunbo{}
	for k := range productList {
		lunbo.Img = productList[k].Img
		lunbo.Url = "/" + productService.TableName() + "/" + productList[k].Name
		lunboList = append(lunboList, lunbo)
	}
	return
}
