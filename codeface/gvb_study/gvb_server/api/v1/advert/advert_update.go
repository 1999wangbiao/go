package advert

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/res"
	"gvb_server/models/system"
)

// AdvertUpdateView 修改数据库中的图片名称(未修改存储的图片路径名称)
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	var cr AdvertRequest

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var advertModel system.AdvertModel
	id := c.Param("id")
	err = global.DB.Take(&advertModel, id).Error
	if err != nil {
		res.FailWithMsg("广告不存在", c)
		return
	}
	err = global.DB.Model(&advertModel).Updates(map[string]any{
		"title":   cr.Title,
		"href":    cr.Href,
		"images":  cr.Images,
		"is_show": cr.IsShow,
	}).Error
	if err != nil {
		res.FailWithMsg("修改广告失败", c)
		return
	}
	res.OKWithMsg("广告名称修改成功", c)
	return
}