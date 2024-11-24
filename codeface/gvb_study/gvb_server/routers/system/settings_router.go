package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type SettingsRouter struct {
}

func (s *SettingsRouter) InitSettingsRouter(Router *gin.Engine) {
	settingsRouter := Router.Group("settings", middleware.JWTAuth())
	settingsApi := api.ApiGroupApp.SystemGroup.SettingsApi
	settingsRouter.GET("settings/:name", settingsApi.SettingsInfoView)
	settingsRouter.PUT("settings/:name", settingsApi.SettingsUpdate)
}
