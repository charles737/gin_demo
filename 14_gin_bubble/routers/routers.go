package routers

import (
	"gin_web_demo/14_gin_bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 找模版文件引用的静态文件
	r.Static("/static", "static")
	// 找模版文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)

		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)

		// 修改某个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateTodo)

		// 删除某个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
