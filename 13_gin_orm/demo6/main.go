package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// gorm demo6 删除操作

// 1. 定义模型
type User6 struct {
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
	db.AutoMigrate(&User6{})

	// 3. 创建
	//u1 := User6{
	//	Name: "老高",
	//	Age: 54,
	//	Active: true,
	//}
	//db.Create(&u1)
	//u2 := User6{
	//	Name: "小茉",
	//	Age: 18,
	//	Active: false,
	//}
	//db.Create(&u2)
	//u3 := User6{
	//	Name: "力气",
	//	Age: 5,
	//	Active: true,
	//}
	//db.Create(&u3)
	//u4 := User6{
	//	Name: "海奥华",
	//	Age: 333333,
	//	Active: false,
	//}
	//db.Create(&u4)

	// 4. 查询
	//var user User6
	//db.First(&user)
	//fmt.Printf("%v\n", user)

	// 5. 删除 软删除 (update users6 set deleted_at=null; 在命令行恢复deleted_at)
	//var u = User6{}
	//u.ID = 1
	//db.Debug().Delete(&u)
	//db.Debug().Where("name=?", "力气").Delete(User6{})
	//db.Debug().Delete(User6{}, "age=?", 18)
	// 物理删除 真正删除数据
	db.Debug().Unscoped().Where("name=?", "海奥华").Delete(User6{})
}
