package menu

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/res"
	"gvb_server/models/system"
)

type MenuNameResponse struct {
	ID    uint   `json:"ID"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

func (MenuApi) MenuNameList(c *gin.Context) {
	var menuNameList []MenuNameResponse
	global.DB.Model(system.MenuModel{}).Select("id", "title", "path").Scan(&menuNameList)
	res.OKWithData(menuNameList, c)
}
