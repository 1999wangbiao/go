package api

import (
	"gvb_server/api/v1/advert"
	"gvb_server/api/v1/images"
	"gvb_server/api/v1/system"
)

type ApiGroup struct {
	SettingsApi system.ApiGroup
	ImagesApi   images.ApiGroup
	AdvertApi   advert.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
