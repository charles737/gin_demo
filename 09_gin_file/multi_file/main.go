package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 多个文件上传

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./index.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("./tmp/%s_%d", file.Filename, index)
			// 上传文件到指定目录
			err := c.SaveUploadedFile(file, dst)
			if err != nil {
				fmt.Printf("save upload file failed, err:%v\n", err)
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("%d files uploaded", len(files)),
			})
		}
	})
	r.Run(":9999")
}
