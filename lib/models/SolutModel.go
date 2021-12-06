package model

type SolutModel struct {
	ID int64 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Fags string `json:"fags" db:"fags"`
	Fajz string `json:"fajz" db:"fajz"`
	Fazc string `json:"fazc" db:"fazc"`
}
