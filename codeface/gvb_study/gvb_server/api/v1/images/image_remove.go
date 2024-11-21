package images

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/res"
	"gvb_server/models/system"
)

func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr system.IDList
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var imageList []system.BannerModel
	rowsAffected := global.DB.Debug().Find(&imageList, cr.IDList).RowsAffected
	if rowsAffected == 0 {
		res.FailWithMsg("文件不存在", c)
		return
	}
	global.DB.Delete(&imageList)
	res.FailWithMsg(fmt.Sprintf("共删除了 %d 张图片", rowsAffected), c)
}
