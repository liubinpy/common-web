package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginApi struct {
}

func (e *LoginApi) Login(c *gin.Context) {
	c.JSON(http.StatusOK, "我是gin")
}
