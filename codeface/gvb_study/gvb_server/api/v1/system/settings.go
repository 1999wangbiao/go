package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/common/response"
)

type SettingsApi struct{}

type SettingsUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 显示全部或者某一项配置信息
// @Tags      配置文件管理
// @Summary	  显示全部或者某一项配置信息
// @Description  显示全部或者某一项配置信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Param     data  body    SettingsUri true "表示多个参数"
// @Success   200   {object}  response.Response{msg=string} "更新成功"
// @Router    /settings/:name [post]
func (s *SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "all":
		response.OKWithData(global.Config, c)
	case "site":
		response.OKWithData(global.Config.SiteInfo, c)
	case "email":
		response.OKWithData(global.Config.Email, c)
	case "jwt":
		response.OKWithData(global.Config.Jwt, c)
	case "qiniu":
		response.OKWithData(global.Config.QiNiu, c)
	case "qq":
		response.OKWithData(global.Config.QQ, c)
	default:
		response.FailWithMsg("没有该参数", c) //
	}
}

// SettingsUpdate 修改某一项配置信息
// @Tags      配置文件管理
// @Summary	  修改某一项配置信息
// @Description  修改某一项配置信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Param     data  body     SettingsUri  true "表示多个参数"
// @Success   200   {object}  response.Response{msg=string} "更新成功"
// @Router    /settings/:name [put]
func (s *SettingsApi) SettingsUpdate(c *gin.Context) {

	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		var cr config.SiteInfo
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.SiteInfo = cr
	case "email":
		var cr config.Email
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.Email = cr
	case "jwt":
		var cr config.Jwt
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.Jwt = cr
	case "qiniu":
		var cr config.QiNiu
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.QiNiu = cr
	case "qq":
		var cr config.QQ
		err := c.ShouldBindJSON(&cr)
		if err != nil {
			response.FailWithCode(response.ArgumentError, c)
			return
		}
		global.Config.QQ = cr
	default:
		response.FailWithMsg("没有该参数", c) //
	}
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		return
	}
	response.OKWith(c)
}
