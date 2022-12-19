package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// setCookie 设置Cookie
func setCookie(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "shan la", Expires: expiration} // 设置Cookie
	http.SetCookie(w, &cookie)
}

// getCookie 获取Cookie
func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("username") // 获取Cookie
	_, err := fmt.Fprint(w, cookie)
	if err != nil {
		log.Fatal("Fprintf Error:", err)
	}
}

func main() {
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe Error:", err)
	}
}
