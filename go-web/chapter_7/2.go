package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct {
	XMLName xml.Name  `xml:"servers"`
	Version string    `xml:"version,attr"`
	Svs     []server1 `xml:"server"`
}

type server1 struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server1{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server1{"Beijing_VPN", "127.0.0.2"})
	// xml编码
	//output, err := xml.Marshal(v)
	output, err := xml.MarshalIndent(v, " ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header)) // 输出xml头
	os.Stdout.Write(output)
}

//<?xml version="1.0" encoding="UTF-8"?>
// <servers version="1">
//     <server>
//         <serverName>Shanghai_VPN</serverName>
//         <serverIP>127.0.0.1</serverIP>
//     </server>
//     <server>
//         <serverName>Beijing_VPN</serverName>
//         <serverIP>127.0.0.2</serverIP>
//     </server>
// </servers>
