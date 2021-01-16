package main

// 模版继承，块模版

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	// 解析模版
	t, err := template.ParseFiles("./index2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	msg := "休士顿"
	// 渲染模版
	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	// 解析模版
	t, err := template.ParseFiles("./home2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	msg := "休士顿"
	// 渲染模版
	t.Execute(w, msg)
}

func index2(w http.ResponseWriter, r *http.Request) {
	// 定义模版（模版继承的方式）
	// 解析模版
	t, err := template.ParseFiles("./template/base.tmpl", "./template/index2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	msg := "芝加哥"
	// 渲染模版
	t.ExecuteTemplate(w, "index2.tmpl", msg)
}

func home2(w http.ResponseWriter, r *http.Request) {
	// 定义模版（模版继承的方式）
	// 解析模版
	t, err := template.ParseFiles("./template/base.tmpl", "./template/home2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	msg := "芝加哥"
	// 渲染模版
	t.ExecuteTemplate(w, "home2.tmpl", msg)
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/index", index)
	http.HandleFunc("/home2", home2)
	http.HandleFunc("/index2", index2)
	err := http.ListenAndServe(":9999",nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v\n", err)
		return
	}
}

