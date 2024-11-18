package global

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvb_server/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
	Router *gin.Engine
)
