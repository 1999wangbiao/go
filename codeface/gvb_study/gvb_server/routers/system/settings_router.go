package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type SettingsRouter struct {
}

func (s *SettingsRouter) InitSettingsRouter(Router *gin.Engine) {
	settingsRouter := Router.Group("api")
	settingsApi := api.ApiGroupApp.SettingsApi
	settingsRouter.GET("settings", settingsApi.SettingsInfoView)
	settingsRouter.PUT("settings", settingsApi.SettingsUpdate)
	settingsRouter.GET("settings_email", settingsApi.SettingsInfoEmailView)
	settingsRouter.PUT("settings_email", settingsApi.SettingsEmailUpdate)
}
