package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/request"
	"gvb_server/models/common/response"
	"gvb_server/models/system"
	"gvb_server/service"
	"gvb_server/service/common"
	"os"
	"path"
	"path/filepath"
)

type ImagesApi struct{}

// ImageListView 查询图片列表
// @Tags      图片管理
// @Summary	  查询图片列表
// @Description  查询图片列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Param     data   body   system.Page  true "表示多个参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /image/imageList [get]
func (i *ImagesApi) ImageListView(c *gin.Context) {
	var cr system.Page
	//绑定第几页参数
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	comList, count, err := common.ComList(system.BannerModel{}, common.Option{
		Page:  cr,
		Debug: false,
	})
	//泛型的使用
	response.OKWithList(comList, count, c)
	return
}

// ImageNameListView 查询图片(只有名字,路径,id)
// @Tags      图片管理
// @Summary	  查询图片(只有名字,路径,id)
// @Description  查询图片(只有名字,路径,id)
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Success   200   {object}  []response.Image
// @Router    /image/imageNames [get]
func (i *ImagesApi) ImageNameListView(c *gin.Context) {
	var imageList []response.Image
	global.DB.Model(system.BannerModel{}).Select("id", "path", "name").Scan(&imageList)
	response.OKWithData(imageList, c)
}

// ImageLoad 上传多个图片
// @Tags      图片管理
// @Summary	  上传多个图片
// @Description  上传多个图片
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Success   200   {object}  response.Response{msg=string}
// @Router    /image/uploadImages [post]
func (i *ImagesApi) ImageLoad(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	fileHeaders := form.File["images"]
	//文件路径是否存在,不存在就创建
	// 将相对路径转换为绝对路径（便于调试和清晰处理）
	absolutePath, err := filepath.Abs(global.Config.UpLoad.Path)
	if err != nil {
		global.Log.Error(err)
		response.FailWithMsg("参数文件路径错误", c)
		return
	}
	_, err = os.Stat(absolutePath)
	if os.IsNotExist(err) {
		// 如果路径不存在，则创建
		err = os.MkdirAll(absolutePath, os.ModePerm) // os.ModePerm 设置目录权限为 0777
		if err != nil {
			global.Log.Error(err)
			response.FailWithMsg("存储文件路径创建失败", c)
			return
		}
	}
	var resList []response.FileUpLoad
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
	response.OKWithData(resList, c)
}

// ImageUpdateView 修改数据库中的图片名称(未修改存储的图片路径名称)
// @Tags      图片管理
// @Summary	  修改数据库中的图片名称(未修改存储的图片路径名称)
// @Description  修改数据库中的图片名称(未修改存储的图片路径名称)
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Param     data  body     request.ImageUpdate true "表示多个参数"
// @Success   200   {object}  response.Response{msg=string}  "修改图片成功"
// @Router    /image/updateImage [put]
func (i *ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr request.ImageUpdate
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	var imageModel system.BannerModel
	err = global.DB.Take(&imageModel, cr.ID).Error
	if err != nil {
		response.FailWithMsg("文件不存在", c)
		return
	}
	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	response.OKWithMsg("图片名称修改成功", c)
	return
}

// ImageRemoveView 删除数据库中的图片
// @Tags      图片管理
// @Summary	  删除数据库中的图片
// @Description  删除数据库中的图片
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Success   200   {object}  response.Response{msg=string}  "修改图片成功"
// @Router    /image/deleteImage [delete]
func (i *ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr system.IDList
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var imageList []system.BannerModel
	rowsAffected := global.DB.Debug().Find(&imageList, cr.IDList).RowsAffected
	if rowsAffected == 0 {
		response.FailWithMsg("文件不存在", c)
		return
	}
	global.DB.Delete(&imageList)
	response.FailWithMsg(fmt.Sprintf("共删除了 %d 张图片", rowsAffected), c)
}
