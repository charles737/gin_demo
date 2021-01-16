package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// orm demo3

// 1. 定义模型
type User1 struct {
	ID int64
	Name string `gorm:"default:'匿名用户'"`
	Gender sql.NullString
	Age int64
}

func main() {
	// 连接MySQL数据库
	db, err := gorm.Open("mysql", "root:mysql123@(127.0.0.1:3306)/go_demo1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	// 使用完后关闭数据库
	defer db.Close()

	// 2. 把模型与数据库中的表对应起来
	db.AutoMigrate(&User1{})

	// 3. 创建
	u1 := User1{Name: "张三", Age: 23} // 在代码层面创建一个 User 对象
	db.Create(&u1) // 在数据库中创建了一条记录
	u2 := User1{Name: "李四", Age: 33} // 在代码层面创建一个 User 对象
	db.Create(&u2) // 在数据库中创建了一条记录
	u3 := User1{Age: 33} // 在代码层面创建一个 User 对象
	db.Create(&u3) // 在数据库中创建了一条记录
	// 当字段类型是sql.NullString时，不能不写，也不能直接写""，必须如下写法
	u4 := User1{Gender: sql.NullString{String: "", Valid: true}, Age: 56}
	db.Create(&u4)
}
