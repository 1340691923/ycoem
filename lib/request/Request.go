package request

import (
	"strconv"

	"github.com/kataras/iris"
	"ycoem/lib/my_error"
)

type Request struct {

}

type CheckConfigStruct struct {
	Code int
	Key  string
}

//检查请求参数
func (this Request) CheckParameter(checkConfig []CheckConfigStruct, ctx iris.Context) (err error) {
	for _, config := range checkConfig {
		if ctx.Method() == "GET"{
			if ctx.Params().Get(config.Key) == "" {
				err = my_error.NewBusiness(ParmasNullError, config.Code)
				return
			}
		}else{
			if ctx.FormValue(config.Key) == "" {
				err = my_error.NewBusiness(ParmasNullError, config.Code)
				return
			}
		}
	}
	return
}

// FormIntDefault 获取Form参数 如果出错则返回默认值
func (this Request) FormIntDefault(ctx iris.Context, key string, def int) int {
	i, err := strconv.Atoi(ctx.FormValue(key))
	if err != nil {
		return def
	}
	return i
}

//获取用户token信息
func (this Request) GetToken(ctx iris.Context) (token string) {
	return ctx.GetHeader("X-Token")
}
