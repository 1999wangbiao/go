package images

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/res"
	"gvb_server/service"
	"gvb_server/service/image_service"
	"os"
	"path"
	"path/filepath"
)

type ImagesApi struct{}

// ImageLoad 上传多个图片
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
	var resList []image_service.FileUpLoadResponse
	for _, fileHeader := range fileHeaders {
		//保存图片
		uploadService := service.ServiceApp.ImageService.ImageUploadService(fileHeader)
		//是否保存成功
		if uploadService.IsSuccess {
			filePath := path.Join(global.Config.UpLoad.Path, fileHeader.Filename)
			err = c.SaveUploadedFile(fileHeader, filePath)
			if err != nil {
				uploadService.Msg = err.Error()
				global.Log.Error(err)
				continue
			}
		}
		resList = append(resList, uploadService)
	}
	res.OKWithData(resList, c)
}
