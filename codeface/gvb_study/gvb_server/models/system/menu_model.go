package system

import (
	"gvb_server/global"
	"gvb_server/models/ctype"
)

type ImageSort struct {
	ImageID uint `json:"image_id"` // 已经上传了的图片，对应file
	Sort    int  `json:"sort"`
}

type MenuModel struct {
	global.GVA_MODEL
	MenuTitle    string        `gorm:"size:32" json:"menu_title"`
	MenuTitleEn  string        `gorm:"size:32" json:"menu_title_en"`
	Icon         string        `gorm:"size:32" json:"icon"`             // 图标名称（前端的图标库名称）
	Path         string        `gorm:"size:32 default:'/'" json:"path"` // 菜单路径
	Slogan       string        `gorm:"size:64" json:"slogan"`
	Abstract     ctype.Array   `gorm:"size:64" json:"abstract"`                                                                   // 简介
	AbstractTime int           `json:"abstract_time"`                                                                             // 简介的切换时间
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;joinReferences:BannerID" json:"banners"` // 菜单的图片列表
	BannerTime   int           `json:"banner_time"`                                                                               // 菜单的切换时间 0表示不切换
	Sort         int           `gorm:"size:10" json:"sort"`                                                                       // 菜单顺序
	ImageSort    []ImageSort   `json:"image_sort" gorm:"-"`                                                                       // 图片具体顺序
}