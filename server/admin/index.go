package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminApi struct {
}

// 登录处理的逻辑
func (e *AdminApi) Index(c *gin.Context) {
	// 获取会话
	c.JSON(http.StatusOK, "我是gin")
}
