package api

import (
	"gvb_server/api/v1/system"
)

type ApiGroup struct {
	SystemGroup system.SystemGroup
}

var ApiGroupApp = new(ApiGroup)
