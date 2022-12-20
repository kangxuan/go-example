package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type ServerSlice struct {
	Servers []Server
}

// 演示Json解码
func main() {
	//var s interface{}
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`

	// 解析到结构体
	var s ServerSlice
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		return
	}
	fmt.Println(s)

	// 解析到interface，会生成一个map
	var f interface{}
	err = json.Unmarshal([]byte(str), &f)
	if err != nil {
		return
	}
	fmt.Println(f)
	// 断言判断
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case int64:
			fmt.Println(k, "is int64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

//{[{Shanghai_VPN 127.0.0.1} {Beijing_VPN 127.0.0.2}]}
//map[servers:[map[serverIP:127.0.0.1 serverName:Shanghai_VPN] map[serverIP:127.0.0.2 serverName:Beijing_VPN]]]
//servers is an array:
//0 map[serverIP:127.0.0.1 serverName:Shanghai_VPN]
//1 map[serverIP:127.0.0.2 serverName:Beijing_VPN]
