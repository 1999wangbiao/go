package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	_ "gvb_server/docs"
	"gvb_server/global"
	"gvb_server/routers/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

func InitRoutes() *gin.Engine {

	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()
	Router.Use(cors.Default()) // 解决跨域问题]
	Router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	systemRouter := RouterGroupApp.System
	systemRouter.InitSettingsRouter(Router)
	systemRouter.InitAdvertRouter(Router)
	systemRouter.InitImagesRouter(Router)
	systemRouter.InitMenuRouter(Router)
	systemRouter.InitUserRouter(Router)
	return Router
}
