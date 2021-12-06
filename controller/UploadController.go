package controller

import (
	"path/filepath"
	"ycoem/config"
	"ycoem/lib/logs"
	"ycoem/lib/service/qiniu"
	"ycoem/lib/service/upload"
	"ycoem/lib/util"

	"github.com/kataras/iris"
	"go.uber.org/zap"
)

// 专门有个上传文件控制器
type UploadController struct {
	BaseController
}

func (this UploadController) ImageAction(ctx iris.Context) {
	const Tag = "ImageAction"
	file, info, err := ctx.FormFile("file")

	if err != nil {
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		this.Error(ctx, err)
		return
	}
	defer file.Close()

	uploadService := upload.UploadService{}
	err = uploadService.UploadImg(info.Filename).Upload(file)

	if err != nil {
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		this.Error(ctx, err)
		return
	}

	qiniuFileName := util.GetUUid() + filepath.Ext(info.Filename)

	qiniuFilePath := "image/" + qiniuFileName

	err = qiniu.Upload(uploadService.Filepath, qiniuFilePath)
	if err != nil {
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, "success", map[string]interface{}{"url": config.VisitImgDir + qiniuFileName})

	logs.Logger.Info("上传成功", zap.String("fname", config.VisitImgDir+qiniuFileName))
	return
}

func (this UploadController) ZipAction(ctx iris.Context) {
	const Tag = "ZipAction"
	file, info, err := ctx.FormFile("file")

	if err != nil {
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		this.Error(ctx, err)
		return
	}
	defer file.Close()

	uploadService := upload.UploadService{}
	err = uploadService.UploadZip(info.Filename).Upload(file)

	if err != nil {
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		this.Error(ctx, err)
		return
	}

	qiniuFileName := util.GetUUid() + filepath.Ext(info.Filename)

	qiniuFilePath := "download/" + qiniuFileName

	err = qiniu.Upload(uploadService.Filepath, qiniuFilePath)
	if err != nil {
		logs.Logger.Error(Tag, zap.String("err.Error()", err.Error()))
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, "success", map[string]interface{}{"url": config.VisitDownloadDir + qiniuFileName})
	logs.Logger.Info("上传成功", zap.String("fname", config.VisitDownloadDir+qiniuFileName))
	return
}
