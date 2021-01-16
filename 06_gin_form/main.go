package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// form表单提交与获取参数

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		// 获取form表单提交的数据
		// 方法一：
		//username := c.PostForm("username")
		//password := c.PostForm("password") // 取到，返回值；取不到，返回空字符串
		// 方法二：
		//username := c.DefaultPostForm("username", "somebody")
		//password := c.DefaultPostForm("xxx", "***")
		// 方法三：
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "不存在"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "***" // 取不到，返回***
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name": username,
			"Password": password,
		})
	})
	r.Run(":9999")
}
