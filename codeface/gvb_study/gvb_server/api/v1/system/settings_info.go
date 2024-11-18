package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/res"
)

type SettingsApi struct{}

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OKWithData(global.Config.SiteInfo, c)
}
