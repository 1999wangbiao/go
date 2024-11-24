package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type ImagesRouter struct {
}

func (s *ImagesRouter) InitImagesRouter(Router *gin.Engine) {
	imagesRouter := Router.Group("image")
	imagesApi := api.ApiGroupApp.SystemGroup.ImagesApi
	imagesRouter.POST("uploadImages", imagesApi.ImageLoad)
	imagesRouter.GET("imageList", imagesApi.ImageListView)
	imagesRouter.GET("imageNames", imagesApi.ImageNameListView)
	imagesRouter.DELETE("deleteImage", imagesApi.ImageRemoveView)
	imagesRouter.POST("updateImage", imagesApi.ImageUpdateView)
}
