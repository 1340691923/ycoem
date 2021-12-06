package login

import (
	"time"

	"github.com/garyburd/redigo/redis"

	"ycoem/lib/db"
	"ycoem/lib/util"
)

type LoginTimeUtil struct {
	StartTime time.Time
}

func HaveWxgamecid(queryMap map[string]interface{}) bool {
	if _, ok := queryMap["wxgamecid"]; ok {
		return true
	}
	return false
}

func IsNewUser(afRows int64) bool {
	return afRows == 1
}

func SetLastLoginTime(userId int64) (err error) {
	zset := "UserLastLogin"
	now := time.Now().Unix() //获取时间戳

	conn := db.Redis.Get()
	defer conn.Close()
	_, err = redis.Bool(conn.Do("ZADD", zset, now, userId))
	return
}

func (this *LoginTimeUtil) In55(startDate, endDate string) bool {
	return this.Diff56Day(startDate) <= 0 && this.Diff56Day(endDate) >= 0
}

func (this *LoginTimeUtil) Diff56Day(diffTime string) int64 {
	diff56Day, _ := time.ParseInLocation(util.TimeFormat, diffTime, time.Local)
	return diff56Day.Unix() - this.StartTime.Unix()
}

func (this *LoginTimeUtil) IsHuiLiuUser(lastLogin string) bool {
	lastLoginTime, _ := time.ParseInLocation(util.TimeFormat, lastLogin, time.Local)
	if this.StartTime.Sub(lastLoginTime).Hours() < 48 {
		return true
	}
	return false
}

func (this *LoginTimeUtil) LostTime() int64 {
	return time.Since(this.StartTime).Milliseconds()
}

func (this *LoginTimeUtil) IsGtThreeDay(created string) (b bool) {
	create_time, _ := time.ParseInLocation(util.TimeFormat, created, time.Local)
	diffDay := this.StartTime.Sub(create_time).Hours() / 24
	if diffDay >= 2 {
		return true
	} else {
		return false
	}
}
