package main

// querystring
// GET 请求 URL ? 后面是 querystring 参数
// key=value格式，多个key=value用 & 连接
// eq: /web?query=邓紫棋&age=22

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器发请求时携带的 query string 参数
		// 第一种查询方式
		name := c.Query("query") // 通过 Query 获取请求中携带的 querystring 参数
		age := c.Query("age")
		// 第二种查询方式
		//name := c.DefaultQuery("query", "nobody") // 取不到就使用默认值
		// 第三种查询方式
		//name, ok := c.GetQuery("query") // 取到返回(值，true)，取不到返回 ("", false)
		//if !ok {
		//	// 取不到
		//	name = "nobody"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age": age,
		})
	})
	r.Run(":9999")
}
