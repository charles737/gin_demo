package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由与路由组

func main() {
	r := gin.Default()
	// 访问"/index"的Get请求会走这一条处理逻辑
	//r.GET("/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"method": "Get",
	//	})
	//})
	//r.POST("/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"method": "POST",
	//	})
	//})
	//r.PUT("/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"method": "PUT",
	//	})
	//})
	//r.DELETE("/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"method": "DELETE",
	//	})
	//})
	// any 请求方法大集合
	r.Any("/index", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
			"method": "GET",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
			"method": "POST",
			})
		case http.MethodPut:
			c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
			})
		case http.MethodDelete:
			c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
			})
		}
	})
	// NoRoute
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "该页面不存在",
		})
	})


	r.Run(":9998")
}
