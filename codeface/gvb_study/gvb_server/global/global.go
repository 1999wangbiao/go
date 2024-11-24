package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"gvb_server/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
	//Router                *gin.Engine
	GvaConcurrencyControl = &singleflight.Group{}
	Redis                 *redis.Client
)

// 定义允许上传图片的文件扩展名白名单
var AllowedImageExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
}
