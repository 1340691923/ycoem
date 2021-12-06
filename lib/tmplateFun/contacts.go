package tmplateFun

import (
	"strings"

	"ycoem/lib/db"
	model "ycoem/lib/models"
	"ycoem/lib/response"
)


func GetContacts()(showContacts response.ShowContacts){
	str:= db.Get(model.ContactsModel)
	showContacts.Address = GetJsonValue(str,"address")
	showContacts.DeptInfos = strings.Split(GetJsonValue(str,"dept_infos"),",")
	return
}
