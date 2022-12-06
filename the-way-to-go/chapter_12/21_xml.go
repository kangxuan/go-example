package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
)

var t, token xml.Token
var err error

func main() {
	// 展示xml的解码和编码
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

	// 解码
	input1 := []byte(input)
	var v interface{}
	err = xml.Unmarshal(input1, &v)
	if err != nil {
		log.Println("解码失败")
		return
	}
	fmt.Println(v)
}
