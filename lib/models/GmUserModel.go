package model

import (
	"go.uber.org/zap"
	"ycoem/lib/db"
	"ycoem/lib/logs"
)

type GmUserModel struct {
	ID       int32  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	RoleId   int32  `json:"role_id" db:"role_id"`
	Realname string `json:"realname" db:"realname"`
}

func (this GmUserModel) Exsit() (b bool) {
	var count int
	err := db.Sqlx.Get(&count, "select count(*) from gm_user where username = ? and password = ? and role_id = ? limit 1;", this.Username, this.Password, this.RoleId)
	if err != nil || count == 0 {
		logs.Logger.Error("err", zap.String("err.Error()", err.Error()))
		return false
	}
	return true
}

func (this GmUserModel) GetUserByUP(username, password string) (gmUser GmUserModel, err error) {
	err = db.Sqlx.Get(&gmUser, "select id,username,password,role_id,realname from gm_user where username = ? and password = ? limit 1;", username, password)
	return
}

func (this GmUserModel) GetUserById() (gmUser GmUserModel, err error) {
	err = db.Sqlx.Get(&gmUser, "select id,username,password,role_id,realname from gm_user where id = ?;", this.ID)
	return
}

func (this GmUserModel) Insert() (id int64, err error) {
	rlt, err := db.Sqlx.Exec("insert into gm_user(username,password,role_id,realname)values(?,?,?,?)", this.Username, this.Password, this.RoleId, this.Realname)
	if err != nil {
		return
	}
	id, _ = rlt.LastInsertId()
	return
}

func (this GmUserModel) Update() (err error) {
	_, err = db.Sqlx.Exec("update gm_user set username = ?,password=?,role_id=?,realname=? where id = ? ;", this.Username, this.Password, this.RoleId, this.Realname, this.ID)
	return
}

func (this GmUserModel) Select() (gmUser []GmUserModel, err error) {
	err = db.Sqlx.Select(&gmUser, "select id,username,password,role_id,realname from gm_user ;")
	return
}

func (this GmUserModel) Delete() (err error) {
	_, err = db.Sqlx.Exec("delete from gm_user where id = ? ;", this.ID)
	return
}
