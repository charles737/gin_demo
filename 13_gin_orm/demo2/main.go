package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// orm demo2 Model

// 定义模型
type User struct {
	gorm.Model // 内嵌 gorm.Model
	Name string
	Age sql.NullInt64 // 零值类型
	Birthday *time.Time
	Email string `gorm:"type:varchar(100);unique_index"`
	Role string `gorm:"size:255"` // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号(member number)唯一且不为空
	Num int `gorm:"AUTO_INCREMENT"` // 设置 Num 为自增类型
	Address string `gorm:"index:addr"` // 给 address 字段创建名为 addr 的索引
	IgnoreMe int `gorm:"-"` // 忽略本字段
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
	db.AutoMigrate(&User{})


}
