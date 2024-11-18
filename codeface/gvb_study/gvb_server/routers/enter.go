package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/routers/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGropuApp = new(RouterGroup)

func InitRoutes() *gin.Engine {

	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()
	Router.Use(cors.Default()) // 解决跨域问题]
	systemRouter := RouterGropuApp.System
	systemRouter.InitSettingsRouter(Router)
	return Router
}
