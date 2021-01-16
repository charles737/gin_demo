package main

import (
	"fmt"
	"net/http"
	"html/template"
)

// 原生方法net/http

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2. 解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err:%v\n", err)
		return
	}
	// 3. 渲染模版
	// 传进去一个字符串
	name := "Linux"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("render template failed, err:%v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("http server start failed, err: %v\n", err)
		return
	}
}