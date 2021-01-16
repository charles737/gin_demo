package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// gorm demo5 查询操作

// 1. 定义模型
type User2 struct {
	gorm.Model
	Name string
	Age int64
}

func main() {
	// 连接MySQL数据库
	db, err := gorm.Open("mysql", "root:mysql123@(127.0.0.1:3306)/go_demo1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// 2. 把模型与数据库中的表对应起来
	db.AutoMigrate(&User2{})

	// 3. 创建
	//u1 := User2{
	//	Name: "老高",
	//	Age: 54,
	//}
	//db.Create(&u1)
	//u2 := User2{
	//	Name: "妙见神",
	//	Age: 18,
	//}
	//db.Create(&u2)

	// 4. 查询
	/*
	new 和 make 的区别：
	new 返回的是基本数据类型的指针
	make 返回的是（map, slice, channel）这些类型（申请什么类型，返回什么类型）
	*/
	// 写法一：
	//var user User2 // 声明模型结构体类型 user
	// 根据主键查询第一条记录 主键必须是数字才能用
	//db.First(&user)
	//user := new(User2)
	//db.First(user)
	//fmt.Printf("user:%v\n", user)
	//
	//fmt.Println()
	//
	//// 查询所有的记录
	//var users []User2
	//db.Debug().Find(&users)
	//fmt.Printf("users:%v\n", users)\

	// FirstOrInit 获取匹配的第一条记录，否则根据给定的条件初始化一个新的对象
	var user User2
	db.FirstOrInit(&user, User2{Name: "小茉"})
	fmt.Printf("%v\n", user)
}
