package advert

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/res"
	"gvb_server/models/system"
)

func (AdvertApi) ImageRemoveView(c *gin.Context) {
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
