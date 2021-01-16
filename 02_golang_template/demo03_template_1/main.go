package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义一个函数kua
	// 要么只有一个返回值，要么有两个返回值，且第二个返回值必须是error类型
	own := func(arg string) (string, error) {
		return arg + "马尔拉戈", nil
	}
	// 定义模版
	t := template.New("template1.tmpl")
	// 告诉模版引擎，现在多了一个自定义函数 own
	t.Funcs(template.FuncMap{
		"own_1": own,
	})
	// 解析模版
	// 和之前的区别在于：先创建了一个带名字的模版对象，然后拿它去链式操作
	_, err := t.ParseFiles("./template1.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	name := "唐纳德"
	// 渲染模版
	t.Execute(w, name)
}

func demoT(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	// 解析模版
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
		return
	}
	name := "唐纳德"
	// 渲染模版
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demoT)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v\n", err)
		return
	}
}

