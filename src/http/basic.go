package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	//http.HandleFunc("/", SayHelloName)
	//http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8000", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func SayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Wrold!") //这个写入到w的是输出到客户端的
}

type CounterHandler struct{}

var counter = 0

func (handler CounterHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	counter = counter + 1
	fmt.Println(w, counter)
}
