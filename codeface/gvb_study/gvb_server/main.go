package main

import (
	"gvb_server/core"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
)

func main() {
	//读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("嘻嘻嘻")
	global.Log.Infoln("哈哈哈")
	//连接数据库
	global.DB = core.InitGrom()
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
