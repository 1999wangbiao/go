package image_service

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models/system"
	"gvb_server/utils"
	"mime/multipart"
	"path"
	"path/filepath"
	"strings"
)

type FileUpLoadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

// ImageUploadService 处理文件上传的方法
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUpLoadResponse) {

	filename := file.Filename
	res.FileName = filename
	//设置图片文件过滤(白名单)
	// 获取文件扩展名并转为小写
	fileExtension := strings.ToLower(filepath.Ext(file.Filename))
	if !global.AllowedImageExtensions[fileExtension] {
		res.Msg = "文件类型不正确!"
		return
	}

	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.UpLoad.Size) {
		res.Msg = fmt.Sprintf("图片大小为%fMB超过设定大小,设定大小为:%dMB!", size, global.Config.UpLoad.Size)
		return
	}
	//获取图片hash
	hash, err := utils.GenerateMD5FromFileHeader(file)
	if err != nil {
		global.Log.Error(err)
		return
	}
	//通过hash查询图片是否已经存在数据库
	var bannerModel system.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", hash).Error
	if err == nil {
		//已经存在数据库
		res.Msg = "该图片已经存在."
		return
	}

	filePath := path.Join(global.Config.UpLoad.Path, file.Filename)
	res.IsSuccess = true
	res.Msg = "上传成功!"
	//插入数据库
	global.DB.Create(&system.BannerModel{
		Path: filePath,
		Hash: hash,
		Name: file.Filename,
	})
	return
}
