package advert

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/common/res"
	"gvb_server/models/system"
	"gvb_server/service/common"
)

// AdvertListView 查询广告列表
func (AdvertApi) AdvertListView(c *gin.Context) {
	//设置接受传递的参数
	/*   type Page struct {
		Page  int    `form:"page"` //第几页
		Key   string `form:"key"`
		Limit int    `form:"limit"` //每一页最多几个
		Sort  string `form:"sort"`
	}*/
	var cr system.Page
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//判断 Refer 是否包含admin ,如果是,就返回全部,不是则返回is_show = true
	/*
		refer := c.GetHeader("Refer")
		is_show := true
		if strings.Contains(referer,"admin"){
		//判断isShow来源
		isShow = false
		}
			//判断refer
				list, count, _ := common.ComList(system.AdvertModel{Is_show:is_show}, common.Option{
					Page: cr,
				})
	*/
	list, count, _ := common.ComList(system.AdvertModel{}, common.Option{
		Page: cr,
	})
	res.OKWithList(list, count, c)
}
