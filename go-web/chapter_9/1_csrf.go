package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

// login 登录逻辑，采用token随机串规避csrf攻击
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		currentTime := time.Now().Unix() // 当前时间戳
		h := md5.New()
		_, err := io.WriteString(h, strconv.FormatInt(currentTime, 10))
		if err != nil {
			return
		}
		token := fmt.Sprintf("%x", h.Sum(nil)) // 生成token

		t, _ := template.ParseFiles("1_csrf_login.html")
		log.Println(t.Execute(w, token))
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("解析Form数据失败")
		}

		token := r.Form.Get("token")
		if token == "" {
			_, err2 := fmt.Fprintln(w, "请勿重复提交")
			if err2 != nil {
				log.Fatal("提交失败")
			}
		}
		log.Println("token: ", token)

		fmt.Println("username: ", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password: ", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) // 输出到客户端
	}
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe Error:", err)
	}
}
