package main

import (
	"fmt"
	"html/template"
)

// 展示template.Must的使用方法
func main() {
	t := template.New("first")
	template.Must(t.Parse(" some static text /* and a comment */"))
	fmt.Println("first 校验通过")

	template.Must(template.New("second").Parse("some static text {{ .Name }}"))
	fmt.Println("second 校验通过")

	tErr := template.New("fail")
	template.Must(tErr.Parse("{{.Name}"))
}
