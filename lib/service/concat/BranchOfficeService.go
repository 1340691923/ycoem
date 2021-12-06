package concat

import (
	model "ycoem/lib/models"
	"ycoem/lib/request"
	"ycoem/lib/service/abstract"
)

type BranchOfficeService struct {

}
func(this *BranchOfficeService) Save(logistics request.BranchOffice)(err error){
	err = abstract.RedisSave(model.BranchOfficeModel,logistics)
	return
}

func(this *BranchOfficeService) Get()(logistics request.BranchOffice,err error){
	err = abstract.RedisGet(model.BranchOfficeModel,&logistics)
	return
}
