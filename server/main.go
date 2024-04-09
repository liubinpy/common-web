package main

import (
	"common-web/server/core"
	"common-web/server/global"
	"common-web/server/initlization"
)

func main() {

	// 初始化配置文件
	core.Viper()

	// 初始化日志
	global.CommonLog = core.InitLogger()

	// 初始化数据库
	global.CommonMysql = initlization.InitMysql()
	initlization.Migrate()

	// 初始化路由
	initlization.WebRouterInit()
}
