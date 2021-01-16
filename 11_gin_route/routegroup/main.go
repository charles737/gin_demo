package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由组

func main() {
	r := gin.Default()
	// 视频的首页和详情页
	// 把公用的前缀提取出来，创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
		})
		videoGroup.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/login"})
		})
		videoGroup.GET("/detail", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/details"})
		})
	}

	r.Run(":9999")
}
