package model

type ServiceAnswerModel struct {
	ID int `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Answer string `json:"answer" db:"answer"`
	ParentID string `json:"parent_id" db:"parent_id"`
	Img string `json:"img" db:"img"`
}

