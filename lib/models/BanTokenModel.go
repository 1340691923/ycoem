package model

import (
	"github.com/garyburd/redigo/redis"

	"ycoem/lib/db"
)

const BanToken = "bantoken_"

func AddBanToken(token string, exptime int64) (b bool, err error) {
	conn := db.Redis.Get()
	defer conn.Close()
	str, err := redis.String(conn.Do("SETEX", BanToken+token, exptime))
	if str == "" {
		b = false
	} else {
		b = true
	}
	return
}

func IsBanToken(token string) (b bool, err error) {
	conn := db.Redis.Get()
	defer conn.Close()
	str, err := redis.String(conn.Do("get", BanToken+token))
	if str == "" {
		b = false
	} else {
		b = true
	}
	return
}
