package main

import (
	"encoding/json"
	"fmt"
)

type Server1 struct {
	ServerName1 string `json:"serverName"` // struct tag
	ServerIP    string
}

type ServerSlice1 struct {
	Servers []Server1
}

// 演示Json编码
func main() {
	var s ServerSlice1
	s.Servers = append(s.Servers, Server1{ServerName1: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server1{ServerName1: "Beijing_VPN", ServerIP: "127.0.0.2"})

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json marshal err: ", err)
	}
	fmt.Println(string(b))
}

//{"Servers":[{"ServerName":"Shanghai_VPN","ServerIP":"127.0.0.1"},{"ServerName":"Beijing_VPN","ServerIP":"127.0.0.2"}]}
