package main

import (
	_ "github.com/swaggo/gin-swagger"
	"gvb_server/core"
	_ "gvb_server/docs"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
)

// @title gvb_server API 文档
// @version 1.0
// @description 这是一个gvb_server API 文档
// @contact.name API Support
// @contact.email developer@example.com
// @host localhost:8080
// @BasePath /
func main() {
	//读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	//router
	router := routers.InitRoutes()
	addr := global.Config.System.Addr()
	global.Log.Info("程序运行在：%s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatal(err)
	}
}
