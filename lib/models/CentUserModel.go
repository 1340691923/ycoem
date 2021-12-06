package model

import (
	"errors"

	"ycoem/lib/db"
)

type CentUserModel struct {
	Id        int64  `json:"id"`
	Chanuser  string `json:"chanuser"`
	Token     string `json:"token"`
	Scene     int    `json:"scene"`
	Created   string `json:"created"`
	LastLogin string `json:"last_login"`
}

//获取注册时间
func (this *CentUserModel) GetCreated() (date string, err error) {
	err = db.Sqlx.Get(&date, "select date(created) from cent_user where id = ?;", this.Id)
	if err != nil {
		return
	}
	return
}

func (this *CentUserModel) GetChanUser() (chanuser string, err error) {
	err = db.Sqlx.Get(&chanuser, "select chanuser from cent_user where id = ?;", this.Id)
	if err != nil {
		return
	}
	return
}

func (this *CentUserModel) CurrDayAddUser() (day int, err error) {
	err = db.Sqlx.Get(&day, "select  count(*) from cent_user where date(created) = date(now());")
	if err != nil {
		return
	}
	return
}

func (this *CentUserModel) GetCreatedAndScene() (err error) {

	err = db.Sqlx.QueryRow("select created,scene,last_login,id from cent_user where  chanuser = ? limit 1;", this.Chanuser).Scan(&this.Created, &this.Scene, &this.LastLogin, &this.Id)

	if err != nil {
		return
	}
	return
}

func (this *CentUserModel) GetLastlogin() (err error) {

	err = db.Sqlx.QueryRow("select last_login from cent_user where id = ? limit 1;", this.Id).Scan(&this.LastLogin)

	if err != nil {
		return
	}
	return
}

func (this *CentUserModel) ToBlack() {
	db.Sqlx.Exec(
		"update cent_user set scene = '0' where id = ?",
		this.Id,
	)
	this.Scene = 0
	return
}

func (this *CentUserModel) GetUidByOpenId() (uid int, err error) {
	sql := " select id from cent_user where chanuser = ? limit 1;"
	db.Sqlx.QueryRow(sql, this.Chanuser).Scan(&uid)
	if uid == 0 {
		err = errors.New("没找到该用户")
		return
	}
	return
}

func (this *CentUserModel) CleanUser() (err error) {
	_, err = db.Sqlx.Exec("delete from cent_user where id = ?", this.Id)
	return
}

func (this *CentUserModel) GetUserByUid() (err error) {
	var count int
	sql := " select count(*) from cent_user where id = ? limit 1;"
	db.Sqlx.QueryRow(sql, this.Id).Scan(&count)
	if count == 0 {
		err = errors.New("没找到该用户")
		return
	}
	return
}

func (this *CentUserModel) GetToken() (err error) {
	sql := " select token from cent_user where id = ? limit 1;"
	db.Sqlx.QueryRow(sql, this.Id).Scan(&this.Token)
	if this.Token == "" {
		err = errors.New("没找到该用户")
		return
	}
	return
}
