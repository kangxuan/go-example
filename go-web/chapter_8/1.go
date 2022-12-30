package main

import (
	"fmt"
	"net"
	"os"
)

// 展示net对IP的支持
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}

//go run 1.go 192.168.1.12
//The address is  192.168.1.12
//go run 1.go 192.168.1
//Invalid address
