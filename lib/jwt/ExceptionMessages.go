package jwt

const (
	ERROR_AUTH_TOKEN = 40006
)

var TOKEN_ERROR = map[int]string{
	ERROR_AUTH_TOKEN: "Token生成失败",
}
