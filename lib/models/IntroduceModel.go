package model

const Introduce  = "Introduce"

type IntroduceModel struct {
	Content string `json:"content"`
	List []string `json:"list"`
	Img string `json:"img"`
}
