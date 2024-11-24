package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type SettingsRouter struct {
}

func (s *SettingsRouter) InitSettingsRouter(Router *gin.Engine) {
	settingsRouter := Router.Group("settings")
	settingsApi := api.ApiGroupApp.SystemGroup.SettingsApi
	settingsRouter.GET("settings/:name", settingsApi.SettingsInfoView)
	settingsRouter.PUT("settings/:name", settingsApi.SettingsUpdate)
}
