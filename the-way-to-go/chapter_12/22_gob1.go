package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y int
	Name string
}

func main() {
	var network bytes.Buffer
	// 编码器
	enc := gob.NewEncoder(&network)
	// 解码器
	dec := gob.NewDecoder(&network)

	// 编码
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	var q Q
	// 解码
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	// 输出数据
	fmt.Printf("%q: {%d,%d}\n", q.Name, q.X, q.Y)
}

//"Pythagoras": {3,4}
