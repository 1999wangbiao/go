package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/routers/advert"
	"gvb_server/routers/image"
	"gvb_server/routers/system"
)

type RouterGroup struct {
	System system.RouterGroup
	Image  image.RouterGroup
	Advert advert.RouterGroup
}

var RouterGropuApp = new(RouterGroup)

func InitRoutes() *gin.Engine {

	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()
	Router.Use(cors.Default()) // 解决跨域问题]
	systemRouter := RouterGropuApp.System
	systemRouter.InitSettingsRouter(Router)
	imageRouter := RouterGropuApp.Image
	imageRouter.InitImagesRouter(Router)
	advertRouter := RouterGropuApp.Advert
	advertRouter.InitAdvertRouter(Router)
	return Router
}
