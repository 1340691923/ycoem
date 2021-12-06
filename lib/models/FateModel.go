package model

type FatdModel struct {
	ID int `json:"id" db:"id"`
	Img string `json:"img" db:"img"`
	Title string `json:"title" db:"title"`
	Describe string `json:"describe" db:"represent"`
	SoluID int `json:"solu_id" db:"solu_id"`
}


