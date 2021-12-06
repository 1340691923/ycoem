package product

import (
	"ycoem/lib/db"
	model "ycoem/lib/models"
)

type ProductNodeService struct {
	model.ProductNodeModel
}

func (this ProductNodeService) ProcessSqlWhere(sqlA db.SelectBuilder) db.SelectBuilder {
	if this.Name != "" {
		sqlA = sqlA.Where(db.Eq{"name":this.Name})
	}
	if this.ID != 0 {
		sqlA = sqlA.Where(db.Eq{"id":this.ID})
	}
	if this.ProductID != 0 {
		sqlA = sqlA.Where(db.Eq{"product_id":this.ProductID})
	}
	return sqlA
}

func (this ProductNodeService) TableName() string {
	return "product_node"
}

func (this ProductNodeService) ProcessSqlInsert(sqlA db.InsertBuilder) db.InsertBuilder {
	setMap := map[string]interface{}{
		"name":this.Name,
		"imgs":this.Imgs,
		"product_id":this.ProductID,
	}
	sqlA = sqlA.SetMap(setMap)
	return sqlA
}

func (this ProductNodeService) ProcessSqlUpdate(id int, sqlA db.UpdateBuilder) db.UpdateBuilder {
	setMap := map[string]interface{}{
		"name":this.Name,
		"imgs":this.Imgs,
		"product_id":this.ProductID,
	}

	sqlA = sqlA.
		Where(db.Eq{"id":id}).
		SetMap(setMap)
	return sqlA
}


