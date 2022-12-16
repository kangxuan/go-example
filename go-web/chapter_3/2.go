package main

import (
	"fmt"
	"net/http"
)

// MyMux 我的复用器
type MyMux struct {
}

// ServeHTTP myMux的ServeHTTP实现
func (m *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" { // 当Path是根目录时
		sayHelloName1(w)
		return
	}

	http.NotFound(w, r)
	return
}

// sayHelloName1 随便写的啥
func sayHelloName1(w http.ResponseWriter) {
	_, err := fmt.Fprintf(w, "Hello myroute!")
	if err != nil {
		return
	}
}

func main() {
	mux := &MyMux{}
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		return
	}
}

// 客户端
// GET http://localhost:9090/
// Hello myroute!

// GET http://localhost:9090/1
// 404 page not found
