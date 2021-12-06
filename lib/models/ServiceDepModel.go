package model

const ServiceDepRedisModel  = "ServiceDepRedisModel"

const ContactsModel  = "ContactsModel"

type ServiceDepModel struct {
	ID int32 `json:"id" db:"id"`
	// Area 地区
	Area string `json:"area" db:"area"`
	// Manager 经理
	Manager string `json:"manager" db:"manager"`
	// Support 技术支持
	Support string `json:"support" db:"support"`
	CreateTime string `json:"create_time" db:"create_time"`
	UpdateTime string `json:"update_time" db:"update_time"`
	// Name 事业部名称
	Name string `json:"name" db:"name"`
}

func(this *ServiceDepModel)TableName()string{
	return "service_dep"
}

