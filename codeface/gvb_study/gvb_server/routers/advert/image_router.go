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
	//advertRouter.GET("list", imagesApi.ImageListView)
	//advertRouter.DELETE("delete", imagesApi.ImageRemoveView)
	//advertRouter.PUT("update", imagesApi.ImageUpdateView)
}
