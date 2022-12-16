package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

// login 实现登录的逻辑
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method) // 获取请求的方法
	if r.Method == "GET" {            // GET请求则是显示登录页面
		t, _ := template.ParseFiles("3_login.html")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm() // 解析form表单数据
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}

		v := url.Values{} // request.Form 是一个 url.Values 类型
		v.Set("name", "Ava")
		fmt.Println("name: ", v.Get("name"))

		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])

		fmt.Println("转义之后")
		fmt.Println("username: ", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password: ", template.HTMLEscapeString(r.Form.Get("password")))

		template.HTMLEscape(w, []byte(r.Form.Get("username"))) // 输出到客户端
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
//name:  Ava
//username:  [shanla]
//password:  [123456]
//转义之后
//username:  shanla
//password:  123456
