package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models/common/request"
	"gvb_server/models/common/response"
	"gvb_server/models/system"
)

type MenuApi struct {
}

// MenuCreateView 添加菜单
// @Tags      菜单管理
// @Summary	  添加菜单
// @Description  添加菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Param     data  body     request.Menu true "菜单信息"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /menu/createMenu [post]
func (m *MenuApi) MenuCreateView(c *gin.Context) {
	fmt.Println("menu")
	var cr request.Menu
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	//重复值判断
	var menuList []system.MenuModel
	count := global.DB.Find(&menuList, "title = ? and path = ? ", cr.Title, cr.Path).RowsAffected
	if count > 0 {
		response.FailWithMsg("重复的菜单", c)
		return
	}
	//创建banner数据入库
	menuModel := system.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMsg("菜单添加失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		response.OKWithMsg("菜单添加失败", c)
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
		response.FailWithMsg("菜单图片添加失败", c)
		return
	}
	response.OKWithData("添加菜单成功", c)
	return
}

// MenuListView 查询全部菜单
// @Tags      菜单管理
// @Summary	  查询全部菜单
// @Description  查询全部的菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Success   200   {object}  []response.Menu{}
// @Router    /menu/menuList [get]
func (m *MenuApi) MenuListView(c *gin.Context) {
	//先查询菜单
	var menuList []system.MenuModel //菜单列表 获得全部菜单
	var menuIDList []uint           //查询菜单的id  获得全部菜单id
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)
	//查询连接表
	var menuBanners []system.MenuBannerModel //菜单和图片连接表  在该表中通过菜单id查询该连接表中的该 id信息
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIDList)
	var menus []response.Menu
	for _, model := range menuList {
		//model 就是一个菜单
		var banners []response.Banners
		for _, banner := range menuBanners {
			if model.ID != banner.MenuID {
				continue
			}
			banners = append(banners, response.Banners{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menus = append(menus, response.Menu{
			MenuModel: model,
			Banners:   banners,
		})
	}
	response.OKWithData(menus, c)
	return
}

// MenuNameList 查询某一个菜单的大致信息(菜单标题,路径)
// @Tags      菜单管理
// @Summary	  查询某一个菜单的大致信息(菜单标题,路径)
// @Description  查询某一个菜单的大致信息(菜单标题,路径)
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Success   200   {object}  response.MenuName{}
// @Router    /menu/menuNameList [get]
func (m *MenuApi) MenuNameList(c *gin.Context) {
	var menuNameList []response.MenuName
	global.DB.Model(system.MenuModel{}).Select("id", "title", "path").Scan(&menuNameList)
	response.OKWithData(menuNameList, c)
}

// MenuRemoveView 删除某一个菜单
// @Tags      菜单管理
// @Summary	  删除某一个菜单
// @Description  删除某一个菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Success   200   {object}  response.Response{msg=string}  "删除菜单成功"
// @Router    /menu/removeMenu [delete]
func (m *MenuApi) MenuRemoveView(c *gin.Context) {
	var cr system.IDList
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var menuList []system.MenuModel
	rowsAffected := global.DB.Debug().Find(&menuList, cr.IDList).RowsAffected
	if rowsAffected == 0 {
		response.FailWithMsg("菜单不存在", c)
		return
	}
	fmt.Println(cr.IDList)
	//这两个表应该一起删除 作为事务进行处理
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		//删除第三张连接表
		err = tx.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error(err)
			return err
		}
		//删除菜单表
		err = tx.Delete(&menuList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		response.FailWithMsg("删除失败", c)
		return
	}
	response.FailWithMsg(fmt.Sprintf("共删除了 %d 个菜单", rowsAffected), c)
}

// MenuUpdateView 更新某一个菜单
// @Tags      菜单管理
// @Summary	  更新某一个菜单
// @Description  更新某一个菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Success   200   {object}  response.Response{msg=string}  "更新菜单成功"
// @Router    /menu/updateMenu [PUT]
func (m *MenuApi) MenuUpdateView(c *gin.Context) {
	var cr request.Menu
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")
	//先把之前的banner清空
	var menuModel system.MenuModel

	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		response.FailWithMsg("菜单不存在", c)
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
			response.FailWithMsg("创建图片失败", c)
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
		response.FailWithMsg("修改菜单失败", c)
		return
	}
	response.OKWithMsg("菜单修改成功", c)
	return
}

// MenuDetailView 菜单详情
// @Tags      菜单管理
// @Summary	  菜单详情
// @Description  查询某一个菜单详情
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Success   200   {object}  response.Menu "图片详情"
// @Router    /menu/menuDetail/{id} [get]
func (m *MenuApi) MenuDetailView(c *gin.Context) {
	//第一步 获取查询的菜单id
	id := c.Param("id")
	//第二步 通过菜单ID查询菜单
	var menuModel system.MenuModel //菜单列表 获得全部菜单
	err := global.DB.Take(&menuModel, id).Error
	if err != nil {
		response.FailWithMsg("菜单不存在", c)
		return
	}
	//第三步 通过菜单id查询 连接表
	var menuBanners []system.MenuBannerModel //菜单和图片连接表  在该表中通过菜单id查询该连接表中的该 id信息
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", id)
	//model 就是一个菜单
	var banners []response.Banners
	// 第四步 通过连接表中的图片id 查询图片
	for _, banner := range menuBanners {
		if menuModel.ID != banner.MenuID {
			continue
		}
		banners = append(banners, response.Banners{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	menuResponse := response.Menu{
		MenuModel: menuModel,
		Banners:   banners,
	}
	response.OKWithData(menuResponse, c)
	return
}
