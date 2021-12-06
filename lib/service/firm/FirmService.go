package firm

import (
	"ycoem/lib/db"
	model "ycoem/lib/models"
	"ycoem/lib/service/abstract"
)

type FirmService struct {
	model.FirmModel
}

func (this FirmService) ProcessSqlWhere(sqlA db.SelectBuilder) db.SelectBuilder {
	if this.Title != "" {
		sqlA = sqlA.Where("title like ?", db.CreateLike(this.Title))
	}
	if this.Name != "" {
		sqlA = sqlA.Where(db.Eq{"name": this.Name})
	}
	if this.WxLink == "not null" {
		sqlA = sqlA.Where(" wx_link != '' ")
	}
	sqlA = sqlA.OrderBy("id desc")
	return sqlA
}

func (this FirmService) TableName() string {
	return "firm"
}

func (this FirmService) ProcessSqlInsert(sqlA db.InsertBuilder) db.InsertBuilder {
	setMap := map[string]interface{}{
		"title":   this.Title,
		"img":     this.Img,
		"wx_link": this.WxLink,
		"content": this.Content,
		"name":    this.Name,
	}
	if this.Img == "" {
		setMap["img"] = ""
	}
	sqlA = sqlA.SetMap(setMap)
	return sqlA
}

func (this FirmService) ProcessSqlUpdate(id int, sqlA db.UpdateBuilder) db.UpdateBuilder {
	setMap := map[string]interface{}{
		"title":   this.Title,
		"img":     this.Img,
		"wx_link": this.WxLink,
		"content": this.Content,
		"name":    this.Name,
	}
	if this.Img == "" {
		setMap["img"] = ""
	}

	sqlA = sqlA.
		Where(db.Eq{"id": id}).
		SetMap(setMap)
	return sqlA
}

func (this *FirmService) Save(logistics model.IntroduceModel) (err error) {

	if logistics.Img != "" {
		logistics.Img = logistics.Img
	}

	err = abstract.RedisSave(model.Introduce, logistics)
	return
}

func (this *FirmService) Get() (logistics model.IntroduceModel, err error) {
	err = abstract.RedisGet(model.Introduce, &logistics)
	return
}
