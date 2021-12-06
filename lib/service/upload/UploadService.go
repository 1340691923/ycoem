package upload

import (
	"io"
	"mime/multipart"
	"os"
	"path"

	"ycoem/config"
)

type UploadService struct {
	Filepath string
}

func(this *UploadService)UploadImg(fname string)*UploadService{
	this.Filepath = path.Join(config.ImgDir,fname)
	return this
}

func(this *UploadService)UploadZip(fname string) *UploadService{
	this.Filepath = path.Join(config.DownloadDir,fname)
	return this
}

func(this *UploadService)Upload(file multipart.File)(err error){
	out, err := os.OpenFile(this.Filepath,os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer out.Close()
	_,err = io.Copy(out, file)
	return
}

func(this *UploadService)Upload2(file io.Reader)(err error){
	out, err := os.OpenFile(this.Filepath,os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer out.Close()
	_,err = io.Copy(out, file)
	return
}
