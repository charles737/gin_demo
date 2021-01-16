package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

// 文件上传

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		// 从请求中读取文件
		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			// 将读取到的文件保存到服务端本地
			// 方式一：
			//dst := fmt.Sprintf("./%s", f.Filename)
			// 方式二：
			dst := path.Join("./", f.Filename)
			err = c.SaveUploadedFile(f, dst)
			if err != nil {
				fmt.Printf("save upload file failed, err:%v\n", err)
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.Run(":9999")
}
