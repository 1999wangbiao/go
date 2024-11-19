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
	settingsRouter.GET("settings/:name", settingsApi.SettingsInfoView)
	settingsRouter.PUT("settings/:name", settingsApi.SettingsUpdate)
}
