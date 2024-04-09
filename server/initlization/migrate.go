package initlization

import (
	"common-web/server/global"
	"common-web/server/model/system"
)

func Migrate() {
	global.CommonMysql.AutoMigrate(
		&system.User{},
	)
}
