package system

// MenuBannerModel 菜单图片关联表
type MenuBannerModel struct {
	MenuID      uint        `json:"menu_id"`
	BannerID    uint        `json:"banner_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID" json:"menu_model"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID" json:"banner_model"`
	Sort        int         `gorm:"size:10" json:"sort"`
}
