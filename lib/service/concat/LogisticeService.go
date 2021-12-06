package concat

import (
	model "ycoem/lib/models"
	"ycoem/lib/request"
	"ycoem/lib/service/abstract"
)

type LogisticeService struct {

}

func(this *LogisticeService ) Save(logistics request.Logistics)(err error){
	err = abstract.RedisSave(model.LogisticsModel,logistics)
	return
}

func(this *LogisticeService ) Get()(logistics request.Logistics,err error){
	err = abstract.RedisGet(model.LogisticsModel,&logistics)
	return
}
