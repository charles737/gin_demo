package main

// index 修改默认标识符
// xss html/template 全部转译

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	// 解析模版
	t, err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	msg := "洛杉矶"
	// 渲染模版
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Printf("execute template failed, err:%v\n", err)
		return
	}
}

func xss(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	// 解析模版
	t, err := template.ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	//str := "<script>alert(123);</script>"
	str := "<a href='charles737.me'>charles737的博客</a>"
	// 渲染模版
	t.Execute(w, str)
}

// 通过自定义函数实现部分转译
func xss1(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	// 解析模版
	// 解析模版之前定义一个自定义的函数safe
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	str1 := "<script>alert(123);</script>"
	str2 := "<a href='https://www.charles737.me'>charles737的博客</a>"
	// 渲染模版
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	http.HandleFunc("/xss1", xss1)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v\n", err)
		return
	}
}