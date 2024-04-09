package global

import (
	"common-web/server/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CommonConfig config.Config
	CommonMysql  *gorm.DB
	CommonLog    *zap.Logger
)
