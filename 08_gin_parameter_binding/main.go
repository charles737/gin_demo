package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin 参数绑定

type UserInfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"pwd"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")

	r.GET("/user", func(c *gin.Context) {
		// 笨方法，变量多了太繁琐
		//Username := c.Query("username")
		//Password := c.Query("password")
		//u := UserInfo{
		//	username: username,
		//	password: password,
		//}
		// gin 参数绑定
		var u UserInfo // 声明一个UserInfo类型的结构体u
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/form", func(c *gin.Context) {
		// gin 参数绑定
		var u UserInfo // 声明一个UserInfo类型的结构体u
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.POST("/json", func(c *gin.Context) {
		// gin 参数绑定
		var u UserInfo // 声明一个UserInfo类型的结构体u
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.Run(":9999")
}
