package system

import (
	"common-web/server/global"
	"common-web/server/model/common/response"
	"common-web/server/model/system/data"
	"common-web/server/model/system/param"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type UserApi struct {
}

// ToLogin 用户登录
func (u *UserApi) ToLogin(ctx *gin.Context) {

	loginInfo := param.LoginParam{}
	err := ctx.ShouldBindJSON(&loginInfo)
	if err != nil {
		response.FailWithMessage("用户名密码错误", ctx)
		return
	}

	user, err := userService.FindUserByAccount(loginInfo.Username)
	if err != nil {
		global.CommonLog.Error("登录错误", zap.String("error", err.Error()))
		response.FailWithMessage("用户名密码错误", ctx)
		return
	}

	if user.Password != loginInfo.Password {
		response.FailWithMessage("用户名密码错误", ctx)
		return
	}

	response.OkWithMessage("登录成功", ctx)
}

var (
	store  = base64Captcha.DefaultMemStore
	driver = base64Captcha.DefaultDriverDigit
)

// GetCaptcha 获取验证码
func (u *UserApi) GetCaptcha(ctx *gin.Context) {
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := c.Generate()
	if err != nil {
		global.CommonLog.Error("生成验证码错误", zap.String("error", err.Error()))
		response.FailWithMessage("生成验证码错误", ctx)
		return
	}
	response.OkWithData(data.CaptchaResponse{
		ID:         id,
		Base64Data: b64s,
	}, ctx)

}

// CaptchaVerify 验证验证码
func (u *UserApi) CaptchaVerify(ctx *gin.Context) {
	verifyInfo := param.CaptchaVerifyParam{}
	err := ctx.ShouldBindJSON(&verifyInfo)
	if err != nil {
		global.CommonLog.Error("验证码bindJSON错误", zap.String("error", err.Error()))
		response.FailWithMessage("验证码输入错误", ctx)
		return
	}

	ok := store.Verify(verifyInfo.ID, verifyInfo.Code, true)
	if ok {
		response.Ok(ctx)
	} else {
		response.FailWithMessage("验证码输入错误", ctx)
	}
}
