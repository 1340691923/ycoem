package util

import (
	"database/sql"
	"github.com/garyburd/redigo/redis"
	"io"
)

const (
	ErrNotError = 0

	ErrNoData       = 96
	ErrParamInvalid = 97
	ErrUnknown      = 99
)

var errMsg = map[int]string{
	ErrNotError:     "成功",
	ErrNoData:       "没有数据",
	ErrParamInvalid: "参数错误",
	ErrUnknown:      "未知错误",
}

// ErrCode 状态信息
type ErrCode = int
type Extra = map[string]interface{}

// WriteErr 输出统一的对外格式
func WriteErr(w io.Writer, code ErrCode, extra ...Extra) (int, error) {
	var d map[string]interface{}
	if extra == nil || extra[0] == nil {
		d = map[string]interface{}{}
	} else {
		d = extra[0]
	}
	d["code"] = code
	var ok bool
	d["msg"], ok = errMsg[code]
	if !ok {
		d["msg"] = "unknown code"
	}
	return WriteJSON(w, d)
}

type Error int

func (e Error) Error() string {
	return errMsg[int(e)]
}

func FilterMysqlNilErr(err error) bool {
	if err != nil && err != sql.ErrNoRows {
		return true
	}
	return false
}

func FilterRedisNilErr(err error) bool {
	if err != nil && err != redis.ErrNil {
		return true
	}
	return false
}
