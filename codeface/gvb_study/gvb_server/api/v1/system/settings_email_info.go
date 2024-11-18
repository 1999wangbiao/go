package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/res"
)

func (SettingsApi) SettingsInfoEmailView(c *gin.Context) {
	res.OKWithData(global.Config.Email, c)
}
