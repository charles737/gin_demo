package controller

import (
	"gin_web_demo/14_gin_bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler (c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

/*
	创建待办事项
 */
func CreateTodo (c *gin.Context) {
	// 前端页面添加待办事项，点击提交，会发送请求到这里
	// 1. 从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo) // 绑定到 todo 数据中
	// 2. 存入数据库
	err := models.CreateATodo(&todo)
	// 3. 返回响应
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 200,
		//	"msg": "success",
		//	"data": todo,
		//})
		c.JSON(http.StatusOK, todo)
	}
}

/*
	查看所有的待办事项
 */
func GetTodoList (c *gin.Context) {
	// 查询 todo 这个表里的所有数据
	todoList, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

/*
	更新待办事项
 */
func UpdateTodo (c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效id",
		})
		return
	}
	todo, err := models.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	c.BindJSON(todo)
	err = models.UpdateATodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

/*
	删除待办事项
 */
func DeleteTodo (c *gin.Context) {
	// 1. 拿到 id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效id",
		})
	}
	// 2. 在数据库中删除 id 对应的数据
	err := models.DeleteATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			id:"deleted",
		})
	}
}