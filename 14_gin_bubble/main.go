package main

import (
	"gin_web_demo/14_gin_bubble/dao"
	"gin_web_demo/14_gin_bubble/models"
	"gin_web_demo/14_gin_bubble/routers"
)


func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}

	defer dao.Close() // 程序退出，关闭数据库连接

	// 模型与数据库关联
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run()
}
