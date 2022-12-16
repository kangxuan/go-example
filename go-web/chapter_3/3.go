package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// login 实现登录的逻辑
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method) // 获取请求的方法
	if r.Method == "GET" {            // GET请求则是显示登录页面
		t, _ := template.ParseFiles("3_login.html")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}

		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/login", login) // 路由分配

	err := http.ListenAndServe(":9091", nil) // 监听端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//method:  GET
//2022/12/16 16:26:53 <nil>
//method:  POST
//username:  [shanla]
//password:  [123456]
