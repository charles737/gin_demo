package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 中间件

func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	//go funcXXX(c.Copy()) // 注意：当在中间件里使用多线程时，为了线程安全，传的是值拷贝
	// 计时
	start := time.Now()
	c.Next()
	cost := time.Since(start)
	fmt.Printf("m1 cost:%v\n", cost)
	fmt.Println("m1 out...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Set("name", "老高") // 在上下文 c 中设置值
	c.Next()
	//c.Abort() // 阻止调用后续的处理函数 这里的阻止是指阻止了后续的所有处理函数
	fmt.Println("m2 out...")
}

func m3(c *gin.Context) {
	fmt.Println("m3 in...")
	// 计时
	start := time.Now()
	c.Next()
	cost := time.Since(start)
	fmt.Printf("m3 cost:%v\n", cost)
	fmt.Println("m3 out...")
}

func main() {
	r := gin.Default()
	r.Use(m1, m2, m3) // 全局注册中间件函数 m1, m2
	r.GET("/index", func(c *gin.Context) {
		name, ok := c.Get("name") // 从上下文中取值（跨中间件取值）
		if !ok {
			name = "匿名用户"
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": name,
		})
	})
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	r.Run(":9998")
}




