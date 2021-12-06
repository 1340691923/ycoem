package model

import (
	"errors"
	"fmt"

	"ycoem/lib/db"
)

type BanSceneModel struct {
	Id     int    `json:"id"`
	Types  string `json:"types"`
	Channo string `json:"channo"`
	Scene  string `json:"scene"`
}

func (this *BanSceneModel) Insert(sceneType, chan_no int, scene string) (err error) {
	_, err = db.Sqlx.Exec("insert into ban_scene(type,scene,chan_no) values(?,?,?)", sceneType, scene, chan_no)
	if err != nil {
		return
	}
	return
}

func (this *BanSceneModel) Delete(id int) (err error) {
	if id == 0 {
		return errors.New("id不能为空")
	}
	_, err = db.Sqlx.Exec("delete from ban_scene where id = ?", id)

	if err != nil {
		return
	}
	return
}

func (this *BanSceneModel) GetInfoById() (sceneType, chan_no, scene string, err error) {
	if err = db.Sqlx.QueryRow("select type,scene,chan_no from ban_scene where id = ? limit 1;", this.Id).Scan(&sceneType, &scene, &chan_no); err != nil {
		return
	}
	return
}

func (this *BanSceneModel) Search(page, limit, sceneType, channo int) (data []interface{}, count int, err error) {
	var sqlWhere string = ""
	if sceneType != 0 {
		sqlWhere = sqlWhere + fmt.Sprintf(" and type = %v", sceneType)
	}
	if channo != 0 {
		sqlWhere = sqlWhere + fmt.Sprintf(" and chan_no = %v", channo)
	}
	rows, err := db.Sqlx.Query(
		"select id,scene,type,chan_no from ban_scene where 1 = 1 "+sqlWhere+" order by id desc limit ?,?",
		(page-1)*limit, limit,
	)
	if err != nil {
		return
	}
	var scene struct {
		Id     int64  `json:"id"`
		Types  int64  `json:"types"`
		Scene  string `json:"scene"`
		ChanNo int64  `json:"chan_no"`
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&scene.Id, &scene.Scene, &scene.Types, &scene.ChanNo); err != nil {
			return
		}
		data = append(data, scene)
	}
	err = db.Sqlx.QueryRow("select count(id) as count from ban_scene where 1 = 1 " + sqlWhere).Scan(&count)

	return
}
