package system

import (
	"gvb_server/global"
)

type BannerModel struct {
	global.GVA_MODEL
	Path string `json:"path"`                 // 图片路径
	Hash string `json:"hash"`                 // 图片hash
	Name string `gorm:"size:128" json:"name"` // 图片名称
}