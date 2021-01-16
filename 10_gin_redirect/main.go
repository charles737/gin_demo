package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// redirect 请求重定向

func main() {
	r := gin.Default()
	// 请求重定向
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com")
	})

	// 路由重定向
	// 方法一：
	//r.GET("/a", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "/b")
	//})
	// 方法二：
	r.GET("/a", func(c *gin.Context) {
		// 跳转到 /b 对应的路由处理函数
		c.Request.URL.Path = "/b" // 把请求的URL修改
		r.HandleContext(c) // 继续后续的处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})

	r.Run(":9999")
}
