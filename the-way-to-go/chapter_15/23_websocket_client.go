package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

func main() {
	// 连接websocket
	ws, err := websocket.Dial("ws://localhost:12345/websocket", "", "http://localhost/")
	// 连接失败panic
	if err != nil {
		panic("Dial: " + err.Error())
	}

	// 从服务端读取数据
	go readFromServer(ws)
	time.Sleep(5 * time.Second)
}

func readFromServer(ws *websocket.Conn) {
	buf := make([]byte, 100)
	for {
		if _, err := ws.Read(buf); err != nil {
			fmt.Printf("%s\n", err.Error())
			break
		}
	}
}
