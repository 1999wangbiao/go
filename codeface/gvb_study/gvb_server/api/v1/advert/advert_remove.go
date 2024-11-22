package advert

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/res"
	"gvb_server/models/system"
)

// AdvertRemoveView 广告删除
// @Tags      广告管理
// @Summary	  广告删除
// @Description  广告删除
// @Param     data  body     system.IDList true "表示多个参数"
// @Router    /advert/remove [delete]
// @Produce   json
// @Success   200   {object}  res.Response{}
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr system.IDList
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var advertList []system.AdvertModel
	rowsAffected := global.DB.Debug().Find(&advertList, cr.IDList).RowsAffected
	if rowsAffected == 0 {
		res.FailWithMsg("广告不存在", c)
		return
	}
	global.DB.Delete(&advertList)
	res.FailWithMsg(fmt.Sprintf("共删除了 %d 张广告", rowsAffected), c)
}
