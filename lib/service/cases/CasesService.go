package cases

import (
	"ycoem/lib/db"
)

type CasesService struct {
	ID int32 `json:"id"`
	// Title 标题
	Title string `json:"title"`
	// Logo 图片
	Logo string `json:"logo"`
	// Desc 描述
	Desc string `json:"desc"`
	// Type 类型
	Type string `json:"type"`
	// Addr 位置
	Addr       string `json:"addr"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func (this CasesService) ProcessSqlWhere(sqlA db.SelectBuilder) db.SelectBuilder {
	if this.Title != "" {
		sqlA = sqlA.Where("title like ?", db.CreateLike(this.Title))
	}
	if this.Addr != "" {
		sqlA = sqlA.Where(db.Eq{"addr": this.Addr})
	}
	if this.Type != "" {
		sqlA = sqlA.Where(db.Eq{"name": this.Type})
	}
	sqlA = sqlA.OrderBy("id desc")
	return sqlA
}

func (this CasesService) ProcessSqlInsert(sqlA db.InsertBuilder) db.InsertBuilder {
	setMap := map[string]interface{}{
		"title":     this.Title,
		"logo":      this.Logo,
		"represent": this.Desc,
		"name":      this.Type,
		"addr":      this.Addr,
	}

	if this.Logo == "" {
		setMap["logo"] = ""
	}

	sqlA = sqlA.SetMap(setMap)

	return sqlA
}

func (this CasesService) ProcessSqlUpdate(id int, sqlA db.UpdateBuilder) db.UpdateBuilder {
	setMap := map[string]interface{}{
		"title":     this.Title,
		"represent": this.Desc,
		"name":      this.Type,
		"addr":      this.Addr,
		"logo":      this.Logo,
	}
	if this.Logo == "" {
		setMap["logo"] = ""
	}
	sqlA = sqlA.
		Where(db.Eq{"id": id}).
		SetMap(setMap)
	return sqlA
}

func (this CasesService) TableName() string {
	return "cases"
}
