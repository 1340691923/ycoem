package solution

import (
	"ycoem/lib/db"
	model "ycoem/lib/models"
)

type FatdService struct {
	Fatd model.FatdModel
}

func (this FatdService) ProcessSqlWhere(sqlA db.SelectBuilder) db.SelectBuilder {
	if this.Fatd.SoluID != 0 {
		sqlA = sqlA.Where(db.Eq{"solu_id": this.Fatd.SoluID})
	}
	sqlA = sqlA.OrderBy("id desc")
	return sqlA
}

func (this FatdService) TableName() string {
	return "fatd"
}

func (this FatdService) ProcessSqlInsert(sqlA db.InsertBuilder) db.InsertBuilder {
	sqlA = sqlA.SetMap(map[string]interface{}{
		"title":     this.Fatd.Title,
		"img":       this.Fatd.Img,
		"represent": this.Fatd.Describe,
		"solu_id":   this.Fatd.SoluID,
	})
	return sqlA
}

func (this FatdService) ProcessSqlUpdate(id int, sqlA db.UpdateBuilder) db.UpdateBuilder {
	setMap := map[string]interface{}{
		"title":     this.Fatd.Title,
		"represent": this.Fatd.Describe,
		"solu_id":   this.Fatd.SoluID,
	}
	if this.Fatd.Img != "" {
		setMap["img"] = this.Fatd.Img
	}

	sqlA = sqlA.
		Where(db.Eq{"id": id}).
		SetMap(setMap)
	return sqlA
}
