package request

const (
	IdNullError       = 100002
	ProvinceNullError = 100003
	PageNullError     = 100004
	LimitNullError    = 100005
	PlatformNullError = 100006
	GameNullError     = 100007
	SdkNoNullError    = 100008
	SdkInfoNullError  = 100009
	VersionNullError  = 100010
	PreNullError      = 100011
	TaskIDNullError   = 100012
	UidNullError      = 100013
	ChannoNullError   = 100014
	PayNoNullError    = 100015
	GameNameNullError = 100016
	ChanuserNullError = 100017
)

var ParmasNullError = map[int]string{
	IdNullError:       "id不能为空！",
	ProvinceNullError: "省份不能为空！",
	PageNullError:     "页码不能为空！",
	LimitNullError:    "每页条数不能为空！",
	PlatformNullError: "平台不能为空！",
	GameNullError:     "游戏id不能为空！",
	SdkNoNullError:    "sdk不能为空！",
	SdkInfoNullError:  "sdk配置不能为空！",
	VersionNullError:  "版本号不能为空！",
	PreNullError:      "域名前缀不能为空！",
	TaskIDNullError:   "taskId不能为空！",
	UidNullError:      "uid不能为空！",
	ChannoNullError:   "渠道号不能为空！",
	PayNoNullError:    "订单号不能为空！",
	GameNameNullError: "游戏名称不能为空！",
	ChanuserNullError: "渠道方用户id不能为空！",
}
