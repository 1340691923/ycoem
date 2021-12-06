package scene

import (
	"github.com/garyburd/redigo/redis"

	"ycoem/lib/db"
)

const (
	NewUserBanRule = "1"
	OldUserBanRule = "2"
)

func DeleteBan(conn redis.Conn, types, channo, value string) (b bool, err error) {
	b, err = redis.Bool(conn.Do("srem", "ban_scene_"+types+"_"+channo, value))
	return
}

func AddBan(conn redis.Conn, types, channo, value string) (b bool, err error) {
	b, err = redis.Bool(conn.Do("sadd", "ban_scene_"+types+"_"+channo, value))
	return
}

func ListBan(types, channo string) (arr []string, err error) {
	conn := db.Redis.Get()
	defer conn.Close()
	arr, err = redis.Strings(conn.Do("smembers", "ban_scene_"+types+"_"+channo))
	return
}
