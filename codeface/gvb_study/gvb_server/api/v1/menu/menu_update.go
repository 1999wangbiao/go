package menu

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/res"
	"gvb_server/models/system"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")
	//先把之前的banner清空
	var menuModel system.MenuModel

	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	global.DB.Model(&menuModel).Association("Banners").Clear()
	//如果选择了banner,那就添加
	if len(cr.ImageSortList) > 0 {
		var bannerList []system.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			bannerList = append(bannerList, system.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err := global.DB.Create(&bannerList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMsg("创建图片失败", c)
			return
		}
	}

	//普通更新
	err = global.DB.Model(&menuModel).Updates(map[string]any{
		"title":         cr.Title,
		"path":          cr.Path,
		"slogan":        cr.Slogan,
		"abstract":      cr.Abstract,
		"abstract_time": cr.AbstractTime,
		"banner_time":   cr.BannerTime,
		"sort":          cr.Sort,
	}).Error
	if err != nil {
		res.FailWithMsg("修改菜单失败", c)
		return
	}
	res.OKWithMsg("菜单修改成功", c)
	return
}
