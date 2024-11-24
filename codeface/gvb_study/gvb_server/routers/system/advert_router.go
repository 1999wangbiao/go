package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type AdvertRouter struct {
}

func (s *AdvertRouter) InitAdvertRouter(Router *gin.Engine) {
	advertRouter := Router.Group("advert")
	advertApi := api.ApiGroupApp.SystemGroup.AdvertApi
	advertRouter.POST("createAdvert", advertApi.AdvertCreateView)
	advertRouter.GET("advertList", advertApi.AdvertListView)
	advertRouter.DELETE("deleteAdvert", advertApi.AdvertRemoveView)
	advertRouter.PUT("updateAdvert", advertApi.AdvertUpdateView)
}
