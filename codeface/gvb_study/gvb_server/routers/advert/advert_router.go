package advert

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type AdvertRouter struct {
}

func (s *AdvertRouter) InitAdvertRouter(Router *gin.Engine) {
	advertRouter := Router.Group("advert")
	advertApi := api.ApiGroupApp.AdvertApi
	advertRouter.POST("create", advertApi.AdvertCreateView)
	advertRouter.GET("list", advertApi.AdvertListView)
	advertRouter.DELETE("delete", advertApi.ImageRemoveView)
	advertRouter.PUT("update", advertApi.AdvertUpdateView)
}
