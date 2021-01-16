package main

import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "start golang!",
	})
}


func main() {
	r := gin.Default() // 返回默认的路由引擎
	// 指定用户使用GET请求访问/hello时，执行sayHello这个函数
	r.GET("/hello", sayHello)

	// demo of GET
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})

	// POST
	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "POST",
		})
	})

	// PUT
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "PUT",
		})
	})

	// DELETE
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "DELETE",
		})
	})
	
	// 启动服务
	r.Run(":9999")
}

