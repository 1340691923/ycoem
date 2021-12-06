package service

import (
	"ycoem/lib/db"
	model "ycoem/lib/models"
)

type ServiceService struct {
	model.ServiceModel
}

func (this ServiceService) ProcessSqlWhere(sqlA db.SelectBuilder) db.SelectBuilder {
	if this.Name != "" {
		sqlA = sqlA.Where(db.Eq{"name":this.Name})
	}
	if this.ID != 0 {
		sqlA = sqlA.Where(db.Eq{"id":this.ID})
	}
	if this.Type != "" {
		sqlA = sqlA.Where(db.Eq{"type":this.Type})
	}
	return sqlA
}

func (this ServiceService) TableName() string {
	return "service"
}

func (this ServiceService) ProcessSqlInsert(sqlA db.InsertBuilder) db.InsertBuilder {
	setMap := map[string]interface{}{
		"name":this.Name,
		"type":this.Type,
	}
	sqlA = sqlA.SetMap(setMap)
	return sqlA
}

func (this ServiceService) ProcessSqlUpdate(id int, sqlA db.UpdateBuilder) db.UpdateBuilder {
	setMap := map[string]interface{}{
		"name":this.Name,
		"type":this.Type,
	}
	sqlA = sqlA.
		Where(db.Eq{"id":id}).
		SetMap(setMap)
	return sqlA
}
