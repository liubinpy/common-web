package global

import (
	"common-web/server/config"
	"gorm.io/gorm"
)

var (
	CommonConfig config.Config
	CommonMysql  *gorm.DB
)
