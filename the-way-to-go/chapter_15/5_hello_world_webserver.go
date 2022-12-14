package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer: " + err.Error())
	}
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	_, err := fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
	if err != nil {
		return
	}
}

//Inside HelloServer handler
//请求
//###
//GET http://localhost:8080/world
//Hello,world
