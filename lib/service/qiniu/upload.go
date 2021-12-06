package qiniu

// 存储相关功能的引入包只有这两个，后面不再赘述
import (
	"context"

	"ycoem/config"
	"ycoem/lib/logs"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

func Upload(filePath string, fileName string) (err error) {
	logs.Logger.Sugar().Infof("七牛云上传", config.QiniuAccessKey, config.QiniuSecretKey, config.QiniuName)
	localFile := filePath
	bucket := config.QiniuName
	key := fileName
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(config.QiniuAccessKey, config.QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{}

	err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		return
	}
	logs.Logger.Sugar().Infof("七牛云上传", ret)
	return
}
