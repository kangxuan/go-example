package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// sayHelloName 处理
func sayHelloName(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // 解析参数，默认不会解析
	if err != nil {
		return
	}

	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["Scheme"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, " "))
	}
	_, err = fmt.Fprintf(w, "hello shanla!") // 写入到w的输出到客户端
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", sayHelloName) // 设置访问路由

	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 启动服务

// 客户端： GET http://localhost:9090/?url_long=111&url_long=222

// 服务端打印
//map[url_long:[111 222]]
//path /
//scheme
//[]
//key: url_long
//val: 111 222
//map[]
//path /favicon.ico
//scheme
//[]
