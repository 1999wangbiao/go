package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/res"
)

type SettingsApi struct{}

type SettingsUri struct {
	Name string `uri:"name"`
}

// 显示全部或者某一项配置信息
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "all":
		res.OKWithData(global.Config, c)
	case "site":
		res.OKWithData(global.Config.SiteInfo, c)
	case "email":
		res.OKWithData(global.Config.Email, c)
	case "jwt":
		res.OKWithData(global.Config.Jwt, c)
	case "qiniu":
		res.OKWithData(global.Config.QiNiu, c)
	case "qq":
		res.OKWithData(global.Config.QQ, c)
	default:
		res.FailWithMsg("没有该参数", c) //
	}
}
