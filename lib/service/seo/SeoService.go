package seo

import (
	model "ycoem/lib/models"
	"ycoem/lib/request"
	"ycoem/lib/service/abstract"
)

type SeoService struct {
}

func (this *SeoService) Save(logistics request.Seo) (err error) {
	err = abstract.RedisSave(model.SeoModel, logistics)
	return
}

func (this *SeoService) Get() (logistics request.Seo, err error) {
	err = abstract.RedisGet(model.SeoModel, &logistics)
	return
}

func (this *SeoService) SetWebinfo(logistics request.Webinfo) (err error) {
	err = abstract.RedisSave(model.WebinfoModel, logistics)
	return
}

func (this *SeoService) GetWebinfo() (logistics request.Webinfo, err error) {
	err = abstract.RedisGet(model.WebinfoModel, &logistics)
	return
}

func (this *SeoService) SetPhoneQQ(logistics request.PhoneQQ) (err error) {
	err = abstract.RedisSave(model.PhoneQQModel, logistics)
	return
}

func (this *SeoService) GetPhoneQQ() (logistics request.PhoneQQ, err error) {
	err = abstract.RedisGet(model.PhoneQQModel, &logistics)
	return
}
