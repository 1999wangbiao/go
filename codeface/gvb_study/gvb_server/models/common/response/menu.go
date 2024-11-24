package response

import (
	"gvb_server/models/system"
)

type Banners struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type Menu struct {
	system.MenuModel
	Banners []Banners `json:"banners"`
}

type MenuName struct {
	ID    uint   `json:"ID"`
	Title string `json:"title"`
	Path  string `json:"path"`
}
