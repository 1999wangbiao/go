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
	//var imageList []system.BannerModel
	////查询的总共条数
	//rowsAffected := global.DB.Select("id").Find(&imageList).RowsAffected
	//fmt.Println(rowsAffected)
	//offset := (cr.Page - 1) * cr.Limit
	//if offset < 0 {
	//	offset = 0
	//}
	//global.DB.Limit(cr.Limit).Offset(offset).Find(&imageList)
	comList, count, err := common.ComList(system.BannerModel{}, common.Option{
		Page:  cr,
		Debug: false,
	})
	//泛型的使用
	res.OKWithList(comList, count, c)
	return
}
