package images

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/res"
	"os"
	"path"
	"path/filepath"
)

type ImagesApi struct{}

type FileUpLoadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

// 上传多个图片
func (ImagesApi) ImageLoad(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	fileHeaders := form.File["images"]
	//文件路径是否存在,不存在就创建
	// 将相对路径转换为绝对路径（便于调试和清晰处理）
	absolutePath, err := filepath.Abs(global.Config.UpLoad.Path)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("参数文件路径错误", c)
		return
	}
	_, err = os.Stat(absolutePath)
	if os.IsNotExist(err) {
		// 如果路径不存在，则创建
		err = os.MkdirAll(absolutePath, os.ModePerm) // os.ModePerm 设置目录权限为 0777
		if err != nil {
			global.Log.Error(err)
			res.FailWithMsg("存储文件路径创建失败", c)
			return
		}
	}
	var resList []FileUpLoadResponse
	for _, fileHeader := range fileHeaders {
		size := float64(fileHeader.Size) / float64(1024*1024)
		if size >= float64(global.Config.UpLoad.Size) {
			resList = append(resList, FileUpLoadResponse{
				FileName:  fileHeader.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小为%fMB超过设定大小,设定大小为:%dMB", size, global.Config.UpLoad.Size),
			})
			continue
		}
		resList = append(resList, FileUpLoadResponse{
			FileName:  fileHeader.Filename,
			IsSuccess: true,
			Msg:       "上传成功",
		})
		filePath := path.Join(global.Config.UpLoad.Path, fileHeader.Filename)
		err := c.SaveUploadedFile(fileHeader, filePath)
		if err != nil {
			global.Log.Error(err)
			continue
		}
	}
	res.OKWithData(resList, c)
}
