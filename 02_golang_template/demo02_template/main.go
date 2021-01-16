package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Gender string
	Age int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 1. 定义模版
	// 2. 解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	// 3. 渲染模版
	// 传进去一个结构体
	u1 := User{
		Name: "张三丰",
		Gender: "男", // 在引擎中使用某个字段或者某个函数的时候，首字母一定要大写
		Age: 64,
	}

	// 传进去一个map
	m1 := map[string]interface{}{
		"name": "蚊香",
		"gender": "女",
		"age": 81,
	}

	// 传进去一个切片
	hobbyList := []string{
		"game1",
		"game2",
		"game3",
	}

	t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
		"hobby": hobbyList,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v\n", err)
		return
	}
}