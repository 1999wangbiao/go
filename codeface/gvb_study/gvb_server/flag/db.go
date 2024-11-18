package flag

import (
	"gvb_server/global"
	"gvb_server/models/system"
)

func MakeMigrations() {
	var err error

	// 设置多对多连接表
	err = global.DB.SetupJoinTable(&system.UserModel{}, "collectsModels", &system.UserCollectModel{})
	//if err != nil {
	//	global.Log.Errorf("数据库迁移失败1: %v", err)
	//	return
	//}
	err = global.DB.SetupJoinTable(&system.MenuModel{}, "banners", &system.MenuBannerModel{})
	//if err != nil {
	//	global.Log.Errorf("数据库迁移失败2: %v", err)
	//	return
	//}

	// 开始迁移数据库
	global.Log.Infof("开始迁移数据库")

	// 确保表创建顺序正确，ArticleModel 应该在 CommentModel 之前创建
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&system.FeedbackModel{},
		&system.BannerModel{},
		&system.AdvertModel{},
		&system.TagModel{},
		&system.CommentModel{},
		&system.ArticleModel{},
		&system.UserModel{},
		&system.LoginDataModel{},
		&system.MessageModel{},
		&system.MenuModel{},
		&system.MenuBannerModel{},
		&system.UserCollectModel{},
	)

	// 错误处理
	if err != nil {
		global.Log.Errorf("数据库迁移失败: %v", err)
		return
	}

	// 迁移成功
	global.Log.Info("数据库迁移成功")
}
