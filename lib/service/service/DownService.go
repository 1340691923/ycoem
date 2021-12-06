package service

import (
	"ycoem/lib/db"
	model "ycoem/lib/models"
)

type DownService struct {
	model.ServiceDownload
}

func (this DownService) ProcessSqlWhere(sqlA db.SelectBuilder) db.SelectBuilder {
	if this.ID != 0 {
		sqlA = sqlA.Where(db.Eq{"id": this.ID})
	}
	if this.Title != "" {
		sqlA = sqlA.Where(db.Eq{"title": this.Title})
	}
	if this.ParentID != "" {
		sqlA = sqlA.Where(db.Eq{"parent_id": this.ParentID})
	}
	return sqlA
}

func (this DownService) TableName() string {
	return "service_download"
}

func (this DownService) ProcessSqlInsert(sqlA db.InsertBuilder) db.InsertBuilder {
	setMap := map[string]interface{}{
		"download_url": this.DownloadURL,
		"title":        this.Title,
		"version":      this.Version,
		"parent_id":    this.ParentID,
	}

	if this.DownloadURL != "" {
		setMap["download_url"] = this.DownloadURL
	}

	sqlA = sqlA.SetMap(setMap)
	return sqlA
}

func (this DownService) ProcessSqlUpdate(id int, sqlA db.UpdateBuilder) db.UpdateBuilder {
	setMap := map[string]interface{}{
		"download_url": this.DownloadURL,
		"title":        this.Title,
		"version":      this.Version,
		"parent_id":    this.ParentID,
	}
	if this.DownloadURL != "" {
		setMap["download_url"] = this.DownloadURL
	}
	sqlA = sqlA.
		Where(db.Eq{"id": id}).
		SetMap(setMap)
	return sqlA
}
