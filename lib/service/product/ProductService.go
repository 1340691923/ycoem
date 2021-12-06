package product

import (
	"ycoem/lib/db"
	model "ycoem/lib/models"
)

type ProductService struct {
	model.ProductModel
}

func (this ProductService) ProcessSqlWhere(sqlA db.SelectBuilder) db.SelectBuilder {
	if this.Name != "" {
		sqlA = sqlA.Where(db.Eq{"name": this.Name})
	}
	sqlA = sqlA.OrderBy("id desc")
	return sqlA
}

func (this ProductService) TableName() string {
	return "product"
}

func (this ProductService) ProcessSqlInsert(sqlA db.InsertBuilder) db.InsertBuilder {

	setMap := map[string]interface{}{
		"name": this.Name,
		"img":  this.Img,
	}

	if this.Img == "" {
		setMap["img"] = ""
	}

	sqlA = sqlA.SetMap(setMap)
	return sqlA
}

func (this ProductService) ProcessSqlUpdate(id int, sqlA db.UpdateBuilder) db.UpdateBuilder {
	return sqlA
}
