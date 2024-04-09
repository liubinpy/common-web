package video

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoApi struct {
}

func (e *VideoApi) VideoIndex(c *gin.Context) {
	username, _ := c.Get("username")
	// 可以获取到login放入session的数据
	c.JSON(http.StatusOK, "我是VideoApi的首页 ： "+username.(string))
}

func (e *VideoApi) GetVideoDetailById(c *gin.Context) {
	username, _ := c.Get("username")
	// 可以获取到login放入session的数据
	param := c.Param("id")
	c.JSON(http.StatusOK, "我是VideoApi的名字,参数:"+param+"  ： "+username.(string))
}
