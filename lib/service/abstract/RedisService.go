package abstract

import (
	"ycoem/lib/db"

	jsoniter "github.com/json-iterator/go"
)

func RedisSave(tableName string,logistics interface{})(err error){
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b,err := json.Marshal(logistics)
	if err!=nil{
		return
	}
	err = db.Set(tableName,string(b))
	return
}

func RedisGet(tableName string,logistics interface{})(err error){
	str := db.Get(tableName)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if str == "" {
		return
	}
	err = json.Unmarshal([]byte(str),logistics)
	if err!=nil{
		return
	}
	return
}
