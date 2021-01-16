package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// gorm demo5 更新操作

// 1. 定义模型
type User3 struct {
	gorm.Model
	Name string
	Age int64
	Active bool
}

func main() {
	// 连接MySQL数据库
	db, err := gorm.Open("mysql", "root:mysql123@(127.0.0.1:3306)/go_demo1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// 2. 把模型与数据库中的表对应起来
	db.AutoMigrate(&User3{})

	// 3. 创建
	//u1 := User3{
	//	Name: "老高",
	//	Age: 54,
	//	Active: true,
	//}
	//db.Create(&u1)
	//u2 := User3{
	//	Name: "小茉",
	//	Age: 18,
	//	Active: false,
	//}
	//db.Create(&u2)

	// 4. 查询
	var user User3
	db.First(&user)
	fmt.Printf("%v\n", user)

	// 4. 更新
	// Save()
	//user.Name = "Zero"
	//user.Age = 28
	//db.Debug().Save(&user) // 默认会修改所有字段

	// Update
	//db.Debug().Model(&user).Update("name", "妙见神")

	// Updates 更新多个属性
	//m1 := map[string]interface{}{
	//	"name": "力气",
	//	"age": 5,
	//	"active": true,
	//}
	//db.Debug().Model(&user).Updates(m1) // m1 列出来的所有字段都会更新
	//db.Debug().Model(&user).Select("age").Update(m1) // 只更新 age 字段
	//db.Debug().Model(&user).Omit("active").Updates(m1) // 排除 active 字段，更新其他字段

	// 让User3表中所有用户的年龄在原有基础上加二
	// 使用SQL表达式更新
	db.Debug().Model(&User3{}).Update("age", gorm.Expr("age+?", 2))

}
