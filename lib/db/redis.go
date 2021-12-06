package db

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"

	jsoniter "github.com/json-iterator/go"
)

var Redis *redis.Pool

// NewRedisPool 新建一个Redis连接池 URL优先
func NewRedisPool(addr, passwd string, db int) *redis.Pool {
	b := strings.HasPrefix(addr, "redis://")
	var dialFunc func() (redis.Conn, error)
	switch {
	case b && passwd == "":
		dialFunc = func() (redis.Conn, error) {
			return redis.DialURL(addr, redis.DialDatabase(db))
		}
	case b && passwd != "":
		dialFunc = func() (redis.Conn, error) {
			return redis.DialURL(addr, redis.DialDatabase(db), redis.DialPassword(passwd))
		}
	case !b && passwd == "":
		dialFunc = func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, redis.DialDatabase(db))
		}
	case !b && passwd != "":
		dialFunc = func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, redis.DialDatabase(db), redis.DialPassword(passwd))
		}
	}

	return &redis.Pool{
		MaxIdle:   10,
		MaxActive: 200,
		Dial: func() (redis.Conn, error) {
			c, err := dialFunc()
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func Set(key, val string) (err error) {
	conn := Redis.Get()
	defer conn.Close()
	_, err = redis.String(conn.Do("set", key, val))
	return
}

func Get(key string) (str string) {
	conn := Redis.Get()
	defer conn.Close()
	str, _ = redis.String(conn.Do("get", key))
	return
}

// BLPOPUnmarshalJSON 从队列阻塞获取json数据并解析
func BLPOPUnmarshalJSON(conn redis.Conn, key string, data interface{}) error {
	b, err := redis.Bytes(conn.Do("BLPOP", key, 0))
	if err != nil {
		return err
	}
	return json.Unmarshal(b, data)
}

// LPUSHMarshalJSON json数据编码并存入队列
func LPUSHMarshalJSON(conn redis.Conn, key string, data interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("LPUSH", key, b)
	return err
}

func GetSyncKey(conn redis.Conn, key string) (string, error) {
	vals, err := redis.Strings(conn.Do("MGET", key, key+"_sync"))
	if err != nil {
		return "", err
	}
	if vals[1] == "" {
		return vals[0], nil
	}
	time.Sleep(200 * time.Millisecond)
	return GetSyncKey(conn, key)
}

func SetSyncKey(conn redis.Conn, key string, fn func(value string) (string, error)) error {
	isSet, err := redis.String(conn.Do("SET", key+"_sync", "token", "PX", "200", "NX"))
	if err != nil {
		return err
	}
	if isSet != "OK" {
		time.Sleep(50 * time.Millisecond)
		return SetSyncKey(conn, key, fn)
	}

	// 处理业务逻辑
	var tmp = make(chan error)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()
				tmp <- fmt.Errorf("unknown error(panic)")
				return
			}
		}()
		s, err := redis.String(conn.Do("GET", key))
		if err != nil && err != redis.ErrNil {
			tmp <- err
			return
		}
		s, err = fn(s)
		if err != nil {
			tmp <- err
			return
		}
		_, err = conn.Do("SET", key, s)
		if err != nil {
			tmp <- err
			return
		}
		tmp <- err
	}()
	t := time.Tick(50 * time.Millisecond)
	for {
		select {
		case err = <-tmp:
			// 处理完成了
			conn.Do("DEL", key+"_sync")
			return err
		case <-t:
			// 刷新锁的时间
			conn.Do("PEXPIRE", key+"_sync", "200")
		}
	}
}
