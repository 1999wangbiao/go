package images

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/common/res"
	"gvb_server/models/system"
	"gvb_server/service/common"
)

func (ImagesApi) ImageListView(c *gin.Context) {
	var cr system.Page
	//绑定第几页参数
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	comList, count, err := common.ComList(system.BannerModel{}, common.Option{
		Page:  cr,
		Debug: false,
	})
	//泛型的使用
	res.OKWithList(comList, count, c)
	return
}
