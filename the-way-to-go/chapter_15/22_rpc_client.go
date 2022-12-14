package main

import (
	"fmt"
	"log"
	"net/rpc"
	"test_15/rpc_objects"
)

const serverAddress = "localhost"

func main() {
	// 连接服务
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}

	args := &rpc_objects.Args{N: 7, M: 8}

	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}

//Args: 7 * 8 = 56
