package initlization

import (
	"common-web/server/core"
	"common-web/server/global"
	"common-web/server/middleware"
	"common-web/server/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func WebRouterInit() {
	//  初始化 gin服务
	rootRouter := gin.New()
	rootRouter.Use(middleware.GinLogger(global.CommonLog),
		middleware.GinRecovery(global.CommonLog, true))

	systemRouter := router.GroupApp.System

	apiGroup := rootRouter.Group("api")
	{
		systemRouter.InitUserRouter(apiGroup)
	}

	// 首页路由
	AdminGroup := rootRouter.Group("admin")
	AdminGroup.Use(middleware.LoginInterceptor())

	// 启动HTTP服务,可以修改端口
	address := fmt.Sprintf(":%d", global.CommonConfig.System.Port)
	server := core.InitServer(address, rootRouter)
	global.CommonLog.Error(server.ListenAndServe().Error())
}
