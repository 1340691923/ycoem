package model

import (
	"go.uber.org/zap"
	"ycoem/lib/db"
	"ycoem/lib/logs"
)

// GmOperater ...
type GmOperaterModel struct {
	ID             int    `json:"id"`
	Operater       string `json:"operater"`
	OperaterID     int    `json:"operater_id"`
	OperaterAction string `json:"operater_action"`
	Created        string `json:"created"`
}

func (this *GmOperaterModel) Insert() {
	_, err := db.Sqlx.Exec("insert into gm_operater (operater,operater_id,operater_action,created) values (?,?,?,now());", this.Operater, this.OperaterID, this.OperaterAction)
	if err != nil {
		logs.Logger.Error("GmOperaterModel", zap.String("err.Error()", err.Error()))
	}
}

