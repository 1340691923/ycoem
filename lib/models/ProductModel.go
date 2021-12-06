package model

type ProductModel struct {
	ID int `json:"id" db:"id"`
	Name string `json:"name"  db:"name"`
	Img string `json:"img" db:"img"`

}

func(this *ProductModel)TableName() string {
	return "product"
}

