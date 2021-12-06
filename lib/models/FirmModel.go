package model

type FirmModel struct {
	ID int `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Img string `json:"img" db:"img"`
	WxLink string `json:"wx_link" db:"wx_link"`
	Content string `json:"content" db:"content"`
	Name string `json:"name" db:"name"`
	CreateTime string `json:"create_time" db:"create_time"`
}
const FirmCacheForIndex  = "FirmCacheForIndex"
