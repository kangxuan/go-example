package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

// MyMux 我的复用器
type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		index(w, r)
		return
	}
	NotFound404(w, r)
}

// NotFound404 404页面
func NotFound404(w http.ResponseWriter, r *http.Request) {
	logrus.Error("页面找不到")
	t, _ := template.ParseFiles("./tmpl/404.tmpl")
	errorInfo := "文件找不到"
	_ = t.Execute(w, errorInfo)
}

// SystemError 系统错误页
func SystemError(w http.ResponseWriter, r *http.Request) {
	logrus.Error("系统错误")
	t, _ := template.ParseFiles("./tmpl/error.tmpl")
	errorInfo := "系统错误"
	_ = t.Execute(w, errorInfo)
}

// index
func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./tmpl/index.tmpl")
	if err != nil {
		fmt.Printf("ParseFiles Error:%v", err)
		return
	}
	// 渲染模板
	err = t.Execute(w, "Shanla")
	if err != nil {
		fmt.Printf("Execute Error:%v", err)
		return
	}
}

func main() {
	myMux := &MyMux{}
	err := http.ListenAndServe(":9090", myMux)
	if err != nil {
		fmt.Printf("ListenAndServe Error:%v", err)
		return
	}
}
