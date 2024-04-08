package video

import (
	"common-web/server/api/video"
	"github.com/gin-gonic/gin"
)

type VideoRouter struct {
}

func (videoRouter VideoRouter) InitVideoRouter(group *gin.RouterGroup) {
	// 帖子路由
	videoApi := video.VideoApi{}
	videoGroup := group.Group("video")
	{
		// 函数封装
		videoGroup.GET("/index", videoApi.VideoIndex)
		videoGroup.GET("/get/:id", videoApi.GetVideoDetailById)
	}
}
