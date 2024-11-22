package advert

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/common/res"
	"gvb_server/models/system"
	"gvb_server/service/common"
)

// AdvertListView 广告列表
// @Tags      广告管理
// @Summary	  广告列表
// @Description  广告列表
// @Param     data  query     system.Page true "表示多个参数"
// @Router    /advert/list [get]
// @Produce   json
// @Success   200   {object}  res.Response{data=res.ListResponse[system.AdvertModel]}
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
