package main

import (
	"encoding/json"
	"fmt"
)

type Server2 struct {
	// - ID不会导出到Json中
	ID         int    `json:"-"`
	ServerName string `json:"serverName"`
	// string选项会将这个字段输出到JSON的时候把值转换
	ServerName2 string `json:"serverName2,string"`
	// 如果带有omitempty选项，name如果该字段值为空，就不会输出到JSON串中
	ServerIP string `json:"serverIP,omitempty"`
}

// 演示Json编码struct tag
func main() {
	var s Server2
	s.ID = 1
	s.ServerName = "host1"
	s.ServerName2 = "host1_copy"
	s.ServerIP = "127.0.0.1"
	b, err := json.Marshal(s)
	if err != nil {
		return
	}
	fmt.Println(string(b))
}

//{"serverName":"host1","serverName2":"\"host1_copy\"","serverIP":"127.0.0.1"}
