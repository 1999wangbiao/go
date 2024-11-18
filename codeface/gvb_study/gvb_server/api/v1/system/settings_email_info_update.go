package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/common/res"
)

func (SettingsApi) SettingsEmailUpdate(c *gin.Context) {
	var cr config.Email
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	fmt.Println("before", global.Config)
	global.Config.Email = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		return
	}
	fmt.Println("later", global.Config)
	res.OKWith(c)
}
