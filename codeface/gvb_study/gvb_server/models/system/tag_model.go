package system

import "gvb_server/global"

type TagModel struct {
	global.GVA_MODEL
	Title    string         `gorm:"size:32" json:"title"`                  // 标签名
	Articles []ArticleModel `gorm:"many2many:article_tag_models" json:"_"` // 标签对应的文章
}
