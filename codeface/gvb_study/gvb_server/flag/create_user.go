package flag

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"gvb_server/models/system"
	"gvb_server/utils/pwd"
)

// CreateUser 创建用户
func CreateUser(permissions string) {
	//创建用户逻辑
	//用户名,昵称 密码 确认密码 邮箱
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)
	fmt.Println("请输入用户名:")
	fmt.Scan(&userName)
	fmt.Println("请输入昵称:")
	fmt.Scan(&nickName)
	for {
		fmt.Println("请输入密码:")
		fmt.Scan(&password)
		fmt.Println("请再次输入密码:")
		fmt.Scan(&rePassword)
		if password == rePassword {
			break
		}
		fmt.Println("两次密码不一致")
	}
	fmt.Println("请输入邮箱:")
	fmt.Scan(&email)

	var userModel system.UserModel
	if global.DB.Take(&userModel, "user_name = ?", userName).Error == nil {
		//	用户已存在
		fmt.Println("用户已存在")
		return
	}

	passwordMD5 := pwd.HashPwd(password)
	userModel = system.UserModel{
		UserName:   userName,
		Password:   passwordMD5,
		NickName:   nickName,
		Email:      email,
		SignStatus: ctype.SignEmail,
		Addr:       "内网注册地址",
		IP:         "127.0.0.1",
		Avatar:     "default.png", // 默认头像
	}
	if permissions == "admin" {
		userModel.Role = ctype.PermissionAdmin
	} else {
		userModel.Role = ctype.PermissionUser
	}
	if global.DB.Create(&userModel).Error != nil {
		fmt.Println("创建用户失败")
	}
	global.Log.Infof(fmt.Sprintf("创建用户成功,用户名:%s,密码:%s", userName, password))
}
