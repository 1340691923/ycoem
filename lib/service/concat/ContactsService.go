package concat

import (
	model "ycoem/lib/models"
	"ycoem/lib/request"
	"ycoem/lib/service/abstract"
)

type ContactsService struct {

}

func(this *ContactsService ) Set(logistics request.Contacts)(err error){
	err = abstract.RedisSave(model.ContactsModel,logistics)
	return
}

func(this *ContactsService) Get()(logistics request.Contacts,err error){
	err = abstract.RedisGet(model.ContactsModel,&logistics)
	return
}
