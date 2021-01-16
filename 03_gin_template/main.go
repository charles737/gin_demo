package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// 静态文件：html页面上用到的 css js 图片

func main() {
	r := gin.Default()
	// 加载静态文件
	r.Static("/statics", "./statics")
	// gin 中给框架添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//r.LoadHTMLFiles("templates/index.tmpl") // 模版解析，加载单个模版文件
	r.LoadHTMLGlob("templates/**/*") // 模版解析，加载templates目录下的所有模版文件 **表示所有目录, *表示所有文件

	r.GET("/posts/index", func(c *gin.Context) {
		// HTTP请求
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ // 模版渲染
			"title": "<a href='https://www.charles737.me'>charles737的博客</a>",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		// HTTP请求
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{ // 模版渲染
			"title": "<a href='https://www.charles737.me'>charles737的博客</a>",
		})
	})

	// 加载从网上下载的模版
	r.GET("home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil) // 如果没有通过define给文件起名字，那么默认是文件名
	})

	r.Run(":9999") // 启动server
}
