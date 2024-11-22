package menu

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/res"
	"gvb_server/models/ctype"
	"gvb_server/models/system"
)

type ImageSort struct {
	ImageID uint `json:"image_id"` //图片id
	Sort    int  `json:"sort"`     //在这个[]ImageSort 排第几
}
type MenuRequest struct {
	MenuTitle     string      `json:"menu_title" binding:"required" msg:"请完善菜单名称"`
	MenuTitleEn   string      `json:"menu_title_en" binding:"required" msg:"请完善菜单英文名称"`
	Slogan        string      `json:"slogan"`
	Abstract      ctype.Array `json:"abstract"`                              // 简介
	AbstractTime  int         `json:"abstract_time"`                         // 简介的切换时间
	BannerTime    int         `json:"banner_time"`                           // 菜单的切换时间 0表示不切换
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号"` // 菜单顺序
	ImageSortList []ImageSort `json:"image_sort_list"`                       //图片的具体顺序
}

// MenuCreateView 添加菜单
// @Tags      菜单管理
// @Summary	  创建菜单
// @Description  创建菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Param     data  body     MenuRequest true "表示多个参数"
// @Success   200   {object}  res.Response{msg=string}  "添加菜单成功"
// @Router    /menu/create [post]
func (MenuApi) MenuCreateView(c *gin.Context) {
	fmt.Println("menu")
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	//创建banner数据入库
	menuModel := system.MenuModel{
		MenuTitle:    cr.MenuTitle,
		MenuTitleEn:  cr.MenuTitleEn,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单添加失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.OKWithMsg("菜单添加失败", c)
		return
	}
	var menuBannerList []system.MenuBannerModel
	for _, sort := range cr.ImageSortList {
		//这里也得判断image_id是否存在这张照片
		menuBannerList = append(menuBannerList, system.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单图片添加失败", c)
		return
	}
	res.OKWithData("添加广告成功", c)
	return
}
