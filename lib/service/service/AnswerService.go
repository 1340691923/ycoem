package service

import (
	"ycoem/lib/db"
	model "ycoem/lib/models"
)

type AnswerService struct {
	model.ServiceAnswerModel
}

func (this AnswerService) ProcessSqlWhere(sqlA db.SelectBuilder) db.SelectBuilder {
	if this.ID != 0 {
		sqlA = sqlA.Where(db.Eq{"id": this.ID})
	}
	if this.ParentID != "" {
		sqlA = sqlA.Where(db.Eq{"parent_id": this.ParentID})
	}
	if this.Title != "" {
		sqlA = sqlA.Where(db.Eq{"title": this.Title})
	}
	return sqlA
}

func (this AnswerService) TableName() string {
	return "service_answer"
}

func (this AnswerService) ProcessSqlInsert(sqlA db.InsertBuilder) db.InsertBuilder {
	setMap := map[string]interface{}{
		"answer":    this.Answer,
		"title":     this.Title,
		"img":       this.Img,
		"parent_id": this.ParentID,
	}

	sqlA = sqlA.SetMap(setMap)
	return sqlA
}

func (this AnswerService) ProcessSqlUpdate(id int, sqlA db.UpdateBuilder) db.UpdateBuilder {
	setMap := map[string]interface{}{
		"answer":    this.Answer,
		"title":     this.Title,
		"img":       this.Img,
		"parent_id": this.ParentID,
	}

	sqlA = sqlA.
		Where(db.Eq{"id": id}).
		SetMap(setMap)
	return sqlA
}
