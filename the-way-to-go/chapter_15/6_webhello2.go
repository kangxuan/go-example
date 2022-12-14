package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/shouthello/", shoutHelloHandler)
	err := http.ListenAndServe("localhost:9999", nil)
	if err != nil {
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 截取/hello/后面的参数
	remPartOfURL := r.URL.Path[len("/hello/"):]
	_, err := fmt.Fprintf(w, "hello %s!", remPartOfURL)
	if err != nil {
		return
	}
}

func shoutHelloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/shouthello/"):]
	_, err := fmt.Fprintf(w, "hello %s!", strings.ToUpper(remPartOfURL))
	if err != nil {
		return
	}
}

//http://localhost:9999/hello/shanla
//hello shanla!

//http://localhost:9999/shouthello/
//hello SHANLA!
