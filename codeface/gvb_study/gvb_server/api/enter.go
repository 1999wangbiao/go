package api

import "gvb_server/api/v1/system"

type ApiGroup struct {
	SettingsApi system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
