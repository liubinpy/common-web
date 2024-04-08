package middleware

import (
	"github.com/gin-gonic/gin"
)

func LoginInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 放行，默认就会放行
	}
}
