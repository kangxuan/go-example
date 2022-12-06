package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

var t xml.Token
var err error

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	// 展示两种解码方式
	input := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	// 生成一个阅读器
	inputReader := strings.NewReader(input)
	// 生成xml的阅读器
	p := xml.NewDecoder(inputReader)

	for t, err = p.Token(); err == nil; t, err = p.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
			}
		case xml.EndElement:
			fmt.Println("End of token")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
		default:
		}
	}

	fmt.Println("--------")

	// 通过Unmarshal解码
	input1 := []byte("<Person><FirstName>La</FirstName><LastName>Shan</LastName></Person>")
	var v interface{}
	err = xml.Unmarshal(input1, &v)
	if err != nil {
		log.Println("解码失败")
		return
	}
	fmt.Println(v)

	fmt.Println("--------")

	// 通过Marshal加密
	person1 := new(Person)
	person1.LastName = "Shan"
	person1.FirstName = "La"
	marshal, err := xml.Marshal(person1)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
}
