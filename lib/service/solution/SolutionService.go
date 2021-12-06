package solution

import (
	"ycoem/lib/db"
	model "ycoem/lib/models"
)

type SolutionService struct {
	model.SolutModel
}

func (this SolutionService) ProcessSqlWhere(sqlA db.SelectBuilder) db.SelectBuilder {
	if this.Name != "" {
		sqlA = sqlA.Where(db.Eq{"name": this.Name})
	}
	sqlA = sqlA.OrderBy("id desc")
	return sqlA
}

func (this SolutionService) TableName() string {
	return "solution"
}

func (this SolutionService) ProcessSqlInsert(sqlA db.InsertBuilder) db.InsertBuilder {
	setMap := map[string]interface{}{
		"name": this.Name,
	}
	setMap["fags"] = this.Fags
	setMap["fajz"] = this.Fajz
	setMap["fazc"] = this.Fazc

	if this.Fags == "" {
		setMap["fags"] = ""
	}
	if this.Fajz == "" {
		setMap["fajz"] = ""
	}
	if this.Fazc == "" {
		setMap["fazc"] = ""
	}

	sqlA = sqlA.SetMap(setMap)

	return sqlA
}

func (this SolutionService) ProcessSqlUpdate(id int, sqlA db.UpdateBuilder) db.UpdateBuilder {
	setMap := map[string]interface{}{
		"name": this.Name,
	}
	setMap["fags"] = this.Fags
	setMap["fajz"] = this.Fajz
	setMap["fazc"] = this.Fazc

	if this.Fags == "" {
		setMap["fags"] = ""
	}
	if this.Fajz == "" {
		setMap["fajz"] = ""
	}
	if this.Fazc == "" {
		setMap["fazc"] = ""
	}

	sqlA = sqlA.
		Where(db.Eq{"id": id}).
		SetMap(setMap)
	return sqlA
}
