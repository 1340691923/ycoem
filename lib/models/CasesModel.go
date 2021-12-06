package model

type CasesModel struct {
	ID int32 `json:"id" db:"id" `
	// Title 标题
	Title string `json:"title" db:"title"`
	// Logo 图片
	Logo string `json:"logo" db:"logo"`
	// Desc 描述
	Desc string `json:"desc" db:"represent"`
	// Type 类型
	Type string `json:"type" db:"name"`
	// Addr 位置
	Addr string `json:"addr" db:"addr"`
	CreateTime string `json:"create_time" db:"create_time"`
	UpdateTime string `json:"update_time" db:"update_time"`
}

const CasesCacheForIndex  = "CasesCacheForIndex"


