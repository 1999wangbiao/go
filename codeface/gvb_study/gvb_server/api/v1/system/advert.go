package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/request"
	"gvb_server/models/common/response"
	"gvb_server/models/system"
	"gvb_server/service/common"
)

type AdvertApi struct {
}

// AdvertCreateView 添加广告
// @Tags      广告管理
// @Summary	  创建广告
// @Description  创建广告
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Param     data  body     request.Advert true "表示多个参数"
// @Success   200   {object}  response.Response{msg=string}  "添加广告成功"
// @Router    /advert/createAdvert [post]
func (a *AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr request.Advert
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	//重复广告判断
	var advert system.AdvertModel
	err = global.DB.Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		response.FailWithMsg("广告已经存在", c)
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
		response.FailWithMsg("添加广告失败", c)
		return
	}
	response.OKWithData("添加广告成功", c)
	return
}

// AdvertRemoveView 广告删除
// @Tags      广告管理
// @Summary	  广告删除
// @Description  广告删除
// @Param     data  body     system.IDList true "表示多个参数"
// @Router    /advert/removeAdvert [delete]
// @Produce   json
// @Success   200   {object}  response.Response{}
func (a *AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr system.IDList
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}

	var advertList []system.AdvertModel
	rowsAffected := global.DB.Debug().Find(&advertList, cr.IDList).RowsAffected
	if rowsAffected == 0 {
		response.FailWithMsg("广告不存在", c)
		return
	}
	global.DB.Delete(&advertList)
	response.FailWithMsg(fmt.Sprintf("共删除了 %d 张广告", rowsAffected), c)
}

// AdvertUpdateView 广告修改
// @Tags      广告管理
// @Summary	  广告修改
// @Description  修改数据库中的图片名称(未修改存储的图片路径名称)
// @Param     data  body     request.Advert true "表示多个参数"
// @Router    /advert/updateAdvert/:id [post]
// @Produce   json
// @Success   200   {object}  response.Response{}
func (a *AdvertApi) AdvertUpdateView(c *gin.Context) {
	var cr request.Advert

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	var advertModel system.AdvertModel
	id := c.Param("id")
	err = global.DB.Take(&advertModel, id).Error
	if err != nil {
		response.FailWithMsg("广告不存在", c)
		return
	}
	err = global.DB.Model(&advertModel).Updates(map[string]any{
		"title":   cr.Title,
		"href":    cr.Href,
		"images":  cr.Images,
		"is_show": cr.IsShow,
	}).Error
	if err != nil {
		response.FailWithMsg("修改广告失败", c)
		return
	}
	response.OKWithMsg("广告名称修改成功", c)
	return
}

// AdvertListView 广告列表
// @Tags      广告管理
// @Summary	  广告列表
// @Description  广告列表
// @Param     data  query     system.Page true "表示多个参数"
// @Router    /advert/advertList [get]
// @Produce   json
// @Success   200   {object}  response.Response{data=response.ListResponse[system.AdvertModel]}
func (a *AdvertApi) AdvertListView(c *gin.Context) {
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
		response.FailWithCode(response.ArgumentError, c)
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
	response.OKWithList(list, count, c)
}
