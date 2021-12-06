package gm_user

const (
	ERROR_AUTH_TOKEN = 40006
	ERROR_AUTH       = 40007
)

var AUTH_ERROR = map[int]string{
	ERROR_AUTH_TOKEN: "Token生成失败",
	ERROR_AUTH:       "用户验证失败",
}
