package model

import (
	"ycoem/lib/db"
	"ycoem/lib/logs"
)

type CentSceneFlagModel struct {
	Flag bool `json:"flag"`
}

//调试模式 开关
func (this *CentSceneFlagModel) GetFlag() bool {
	err := db.Sqlx.QueryRow("select flag from cent_scene_flag limit 1;").Scan(&this.Flag)
	if err != nil {
		logs.Logger.Sugar().Errorf("get flag err", err, this.Flag)
		return false
	}
	return this.Flag
}

func (this *CentSceneFlagModel) SetFlag() error {
	_, err := db.Sqlx.Exec("update cent_scene_flag set flag = ? ;", this.Flag)
	if err != nil {
		return err
	}
	return err
}
