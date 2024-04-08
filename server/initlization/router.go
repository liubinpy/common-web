package initlization

import (
	"common-web/server/api/login"
	"common-web/server/middleware"
	"common-web/server/router/video"
	"fmt"
	"github.com/gin-gonic/gin"
)

func WebRouterInit() {
	//  初始化 gin服务
	rootRouter := gin.Default()

	videoRouter := video.VideoRouter{}
	loginApi := login.LoginApi{}

	// 登录路由
	rootRouter.GET("/login", loginApi.Login)

	// 首页路由
	AdminGroup := rootRouter.Group("admin")
	AdminGroup.Use(middleware.LoginInterceptor())
	{
		videoRouter.InitVideoRouter(AdminGroup)
	}
	// 启动HTTP服务,可以修改端口
	address := fmt.Sprintf(":%d", 8088)
	rootRouter.Run(address)
}
