package images

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/res"
	"gvb_server/models/system"
)

type ImageUpdateView struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择文件id"`
	Name string `json:"name" binding:"required" msg:"请输入文件名称"`
}

// ImageUpdateView 修改数据库中的图片名称(未修改存储的图片路径名称)
func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr ImageUpdateView
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var imageModel system.BannerModel
	err = global.DB.Take(&imageModel, cr.ID).Error
	if err != nil {
		res.FailWithMsg("文件不存在", c)
		return
	}
	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OKWithMsg("图片名称修改成功", c)
	return
}
