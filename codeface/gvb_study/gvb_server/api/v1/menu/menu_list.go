package menu

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/res"
	"gvb_server/models/system"
)

type Banners struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	system.MenuModel
	Banners []Banners `json:"banners"`
}

func (MenuApi) MenuListView(c *gin.Context) {
	//先查询菜单
	fmt.Println("list")
	var menuList []system.MenuModel //菜单列表 获得全部菜单
	var menuIDList []uint           //查询菜单的id  获得全部菜单id
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)
	//查询连接表
	var menuBanners []system.MenuBannerModel //菜单和图片连接表  在该表中通过菜单id查询该连接表中的该 id信息
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIDList)
	var menus []MenuResponse
	for _, model := range menuList {
		//model 就是一个菜单
		var banners []Banners
		for _, banner := range menuBanners {
			if model.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banners{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: model,
			Banners:   banners,
		})

	}
	res.OKWithData(menus, c)
	return
}
