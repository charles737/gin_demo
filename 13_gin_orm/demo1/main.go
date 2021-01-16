package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// orm demo1

// UserInfo --> 数据表
type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}

func main() {
	// 连接MySQL数据库
	db, err := gorm.Open("mysql", "root:mysql123@(127.0.0.1:3306)/go_demo1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	// 使用完后关闭数据库
	defer db.Close()

	// 创建表 自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(UserInfo{})

	// 创建数据行
	//u1 := UserInfo{
	//	ID: 1,
	//	Name: "刘德华",
	//	Gender: "男",
	//	Hobby: "篮球",
	//}
	//db.Create(&u1)

	// 查询
	var u UserInfo
	//db.First(&u) // 查询表中第一条数据，保存到 u 中
	//fmt.Printf("u:%#v\n", u)

	// 更新
	//db.Model(&u).Update("hobby", "足球")

	// 删除
	db.Delete(&u)
}
