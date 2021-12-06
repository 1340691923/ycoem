package gm_user

import (
	"ycoem/lib/jwt"
	model "ycoem/lib/models"
)

type GmUserService struct {
}

func (this GmUserService) CheckLogin(username, password string) (token string, err error) {
	var model2 model.GmUserModel
	gmUser, err := model2.GetUserByUP(username, password)
	if err != nil {
		return
	}
	token, err = jwt.GenerateToken(gmUser)
	if err != nil {
		return
	}
	return
}

func (this GmUserService) GetRoleInfo(roleId int32) (gminfo model.GmRoleModel, err error) {
	var model2 model.GmRoleModel
	gminfo, err = model2.GetById(int(roleId))
	if err != nil {
		return
	}
	return
}

func (this GmUserService) IsExitUser(claims *jwt.Claims) bool {
	var model2 model.GmUserModel
	model2.Username = claims.Username
	model2.Password = claims.Password
	model2.RoleId = claims.RoleId
	return model2.Exsit()
}
