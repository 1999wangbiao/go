package advert

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/res"
	"gvb_server/models/system"
)

type AdvertApi struct{}

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题"`       // 广告标题 唯一
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法"`   // 广告链接
	Images string `json:"images" binding:"required,url" msg:"图片地址非法"` // 图片
	IsShow bool   `json:"is_show" binding:"required" msg:"请选择是否展示"`   // 是否显示
}

func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//重复广告判断
	var advert system.AdvertModel
	err = global.DB.Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMsg("广告已经存在", c)
		return
	}
	err = global.DB.Create(&system.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("添加广告失败", c)
		return
	}
	res.OKWithData("添加广告成功", c)
	return
}
