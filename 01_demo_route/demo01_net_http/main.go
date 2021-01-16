//package main
//
//import (
//	"fmt"
//	"io/ioutil"
//	"net/http"
//)
//
//func sayHello(w http.ResponseWriter, r *http.Request) {
//	b, _ := ioutil.ReadFile("./hello.html")
//	_, _ = fmt.Fprintln(w, string(b))
//}
//
//func main() {
//	http.HandleFunc("/hello", sayHello)
//	err := http.ListenAndServe(":9091", nil)
//	if err != nil {
//		fmt.Printf("http server failed, err:%v\n", err)
//		return
//	}
//}
