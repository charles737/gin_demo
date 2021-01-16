package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var (
	DB *gorm.DB
)

func InitMySQL()(err error){
	dsn := "root:mysql123@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn) // 不要 := （重复声明了变量）
	if err != nil {
		return
	}
	err = DB.DB().Ping() // 测试连通性 ping通时，返回new()
	return
}

func Close(){
	DB.Close()
}