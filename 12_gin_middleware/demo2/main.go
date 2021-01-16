package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由组 中间件

func outMiddleware(doCheck bool) gin.HandlerFunc {
	// 连接数据库
	// 其他准备操作
	return func(c *gin.Context) {
		if doCheck {
			// 存放具体的逻辑
			// 是否登录判断
			//if 是登录用户 {
			//	c.Next()
			//}
			//else {
			//	c.Abort()
			//}
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()
	// 路由组注册中间件方法一：
	videoGroup := r.Group("/video", outMiddleware(true))
	{
		videoGroup.GET("/details", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": " video details",
			})
		})
		videoGroup.GET("/order", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "video order",
			})
		})
	}
	// 路由组注册中间件方法二：
	userGroup := r.Group("user")
	userGroup.Use(outMiddleware(false))
	{
		userGroup.GET("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": " user information",
			})
		})
		videoGroup.GET("/secure", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "user secure",
			})
		})
	}
}
