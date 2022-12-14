package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

func server(ws *websocket.Conn) {
	fmt.Printf("New Connection\n")
	buf := make([]byte, 100)
	for {
		if _, err := ws.ReadMessage(); err != nil {
			fmt.Printf("%s", err.Error())
		}
	}
	fmt.Printf(" => closing connection\n")
	ws.Close()
}

func main() {
	// 处理websocket路由
	http.Handle("/websocket", websocket.Handler(server))
	// 监听12345
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
