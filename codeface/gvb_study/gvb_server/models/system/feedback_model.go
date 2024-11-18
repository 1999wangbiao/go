package system

import "gvb_server/global"

// 用户反馈表

type FeedbackModel struct {
	global.GVA_MODEL
	Email        string `gorm:"size:128" json:"email"`
	Content      string `gorm:"type:longtext" json:"content"`
	ApplyContent string `gorm:"type:longtext" json:"apply_content"` // 回复的内容
	IsApply      bool   `gorm:"default:false" json:"is_apply"`      // 是否回复
}
