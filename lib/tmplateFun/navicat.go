package tmplateFun

import (
	"strings"

	"ycoem/lib/db"
	model "ycoem/lib/models"

	jsoniter "github.com/json-iterator/go"
)

func GetNavicat()(obj interface{})  {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	str:= db.Get(model.NavbarModel)
	json.Unmarshal([]byte(str),&obj)
	return
}
func DeleteLine(str string)string {
	//首页
	if str == "/" {
		return "/"
	}
	return strings.Replace(str,"/","-", -1 )
}
