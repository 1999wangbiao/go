package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/common/res"
)

// 修改某一项配置信息
func (SettingsApi) SettingsUpdate(c *gin.Context) {

	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		var cr config.SiteInfo
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.SiteInfo = cr
	case "email":
		var cr config.Email
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Email = cr
	case "jwt":
		var cr config.Jwt
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Jwt = cr
	case "qiniu":
		var cr config.QiNiu
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QiNiu = cr
	case "qq":
		var cr config.QQ
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QQ = cr
	default:
		res.FailWithMsg("没有该参数", c) //
	}
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		return
	}
	res.OKWith(c)
}
