package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	// 组装数据
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	// 编码
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s\n", js)

	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	// 关闭文件
	defer func(src *os.File) {
		err := src.Close()
		if err != nil {
			return
		}
	}(file)

	// 生成file的jsonEncoder
	enc := json.NewEncoder(file)
	// 将vc的编码写入io.Writer w中
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}

	// 解码
	var f interface{}
	//b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	err = json.Unmarshal(js, &f)
	if err != nil {
		log.Println("解码失败")
		return
	}
	fmt.Printf("JSON UnFormat: %v\n", f)

	// 从文件中读取数据并解码
	file1, _ := os.Open("b.json")
	dec := json.NewDecoder(file1)
	var v interface{}
	err = dec.Decode(&v)
	if err != nil {
		log.Println("Error in decoding json")
	}
	fmt.Printf("JSON UnFormat: %v\n", v)
}
