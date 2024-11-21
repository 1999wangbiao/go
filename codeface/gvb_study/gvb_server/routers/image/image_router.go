package image

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type ImagesRouter struct {
}

func (s *ImagesRouter) InitImagesRouter(Router *gin.Engine) {
	imagesRouter := Router.Group("image")
	imagesApi := api.ApiGroupApp.ImagesApi
	imagesRouter.POST("upload", imagesApi.ImageLoad)
	imagesRouter.GET("list", imagesApi.ImageListView)
	imagesRouter.DELETE("delete", imagesApi.ImageRemoveView)
	imagesRouter.PUT("update", imagesApi.ImageUpdateView)
}
