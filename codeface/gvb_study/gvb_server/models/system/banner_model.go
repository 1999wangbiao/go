package system

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"os"
)

type BannerModel struct {
	global.GVA_MODEL
	Path      string          `json:"path"`                        // 图片路径
	Hash      string          `json:"hash"`                        // 图片hash
	Name      string          `gorm:"size:128" json:"name"`        // 图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` //图片存储本地还是七牛或者其他
}

func (b BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		//如果是本地图片.删除本地图片
		err := os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return
}
