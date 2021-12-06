package middleware

const (
	INVALID_PARAMS                 int = 40001
	ERROR_AUTH_CHECK_TOKEN_FAIL    int = 40002
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT int = 40003
)

var TOKEN_ERROR = map[int]string{
	INVALID_PARAMS:                 "Token不能为空",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
}
