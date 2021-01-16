package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		// 方法一：使用map 自己拼接 json
		//data := map[string]interface{}{
		//	"name": "唐纳德",
		//	"age": 66,
		//	"message": "good night!",
		//}
		data := gin.H{
			"name": "唐纳德",
			"age": 77,
			"message": "good morning!",
		}
		c.JSON(http.StatusOK, data)
	})
	// 方法二：使用结构体
	// 为了读取到，这里必须首字母大写
	// 使用 tag 来对结构体字段灵活化操作
	type msg struct{
		Name string
		Age int
		Message string `json:"message"`
	}
	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			"唐纳德",
			88,
			"good afternoon!",
		}
		c.JSON(http.StatusOK, data) // json 的序列化
	})
	r.Run(":9999")
}
