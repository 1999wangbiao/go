package flag

import (
	sys_flag "flag"
	"gvb_server/global"
)

type Option struct {
	DB   bool
	User string // -u admin  -u user
}

// Parse 解析命令行参数
func Parse() Option {

	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	//解析命令行参数写入注册的flag里
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) bool {
	if option.DB {
		global.Log.Info("停止web项目")
		return true
	}
	if option.User == "user" || option.User == "admin" {
		global.Log.Info("停止web项目")
		return true
	}
	return false
}

func SwitchOption(option Option) {
	if option.DB {
		MakeMigrations()
		return
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}
	sys_flag.Usage()
}
