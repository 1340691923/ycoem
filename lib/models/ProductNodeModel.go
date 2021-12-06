package model

type ProductNodeModel struct {
	ID int `json:"id" db:"id"`
	CreateTime string `json:"create_time" db:"create_time"`
	UpdateTime string `json:"update_time" db:"update_time"`
	ProductID int `json:"product_id" db:"product_id"`
	Imgs string `json:"imgs" db:"imgs"`
	Name string `json:"name" db:"name"`
}
