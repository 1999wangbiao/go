package request

import "gvb_server/models/ctype"

type ImageSort struct {
	ImageID uint `json:"image_id"` //图片id
	Sort    int  `json:"sort"`     //在这个[]ImageSort 排第几
}
type Menu struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径/别名"`
	Slogan        string      `json:"slogan"`
	Abstract      ctype.Array `json:"abstract"`                              // 简介
	AbstractTime  int         `json:"abstract_time"`                         // 简介的切换时间
	BannerTime    int         `json:"banner_time"`                           // 菜单的切换时间 0表示不切换
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号"` // 菜单顺序
	ImageSortList []ImageSort `json:"image_sort_list"`                       //图片的具体顺序
}
