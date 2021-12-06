package config

import (
	"io/ioutil"
	"os"
	"path"

	"ycoem/lib/db"
	"ycoem/lib/logs"

	jsoniter "github.com/json-iterator/go"
)

var (
	Split            = "/"
	AppDir           = ""
	ViewDir          = ""
	LogDir           = ""
	ImgDir           = ""
	VisitImgDir      = ""
	Domain           = ""
	DownloadDir      = ""
	VisitDownloadDir = ""
	Port             = ""
	QiniuAccessKey   = ""
	QiniuSecretKey   = ""
	QiniuName        = ""
)

func Init() {
	AppDir, _ = os.Getwd()
	configMap := getConfig()
	Domain = configMap["pre_domain"].(string)
	Port = configMap["port"].(string)
	setDownloadDir()
	setImgDir()
	setViewDir()
	setLogDir()

	logs.InitLog(LogDir)

	db.Sqlx = db.NewMySQL(configMap["my_sql_source"].(string), 10, 50)
	db.Redis = db.NewRedisPool(configMap["redis_source"].(string), "", int(configMap["redis_db"].(float64)))
	QiniuAccessKey = configMap["qiniu_access_key"].(string)
	QiniuSecretKey = configMap["qiniu_secret_key"].(string)
	QiniuName = configMap["qiniu_name"].(string)
}

func setViewDir() {
	ViewDir = path.Join(AppDir, "view"+Split+"template")
}

func setImgDir() {
	ImgDir = path.Join(AppDir, "file"+Split+"img")
	VisitImgDir = Domain + Split + "image" + Split
}

func setDownloadDir() {
	DownloadDir = path.Join(AppDir, "file"+Split+"download")
	VisitDownloadDir = Domain + Split + "download" + Split
}

func setLogDir() {
	LogDir = path.Join(AppDir, "logs")
}

func getConfig() (configMap map[string]interface{}) {
	configMap = map[string]interface{}{}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("config.json 配置文件没找到:" + err.Error())
	}
	err = json.Unmarshal(b, &configMap)
	if err != nil {
		panic("config.json 解析错误:" + err.Error())
	}
	return
}
