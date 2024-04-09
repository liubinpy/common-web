package system

import (
	v1 "common-web/server/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	systemApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("/toLogin", systemApi.ToLogin)
		userRouter.GET("/getCaptcha", systemApi.GetCaptcha)
		userRouter.POST("/captchaVerify", systemApi.CaptchaVerify)
	}
}
