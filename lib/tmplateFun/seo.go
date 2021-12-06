package tmplateFun

import (
	"ycoem/lib/db"
	model "ycoem/lib/models"
)


func GetSeoObj() string  {
	str:= db.Get(model.SeoModel)
	return str
}

func GetPhoneObj()string {
	str := db.Get(model.PhoneQQModel)
	return str
}

func GetWebInfo()string  {
	str := db.Get(model.WebinfoModel)
	return str
}
