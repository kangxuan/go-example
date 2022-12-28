package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	email    string
}

func main() {
	t := template.New("template name 1")
	t, _ = t.Parse("hello {{.UserName}}, {{.email}}")
	p := Person{UserName: "Shan la", email: "999@qq.com"}
	_ = t.Execute(os.Stdout, p)
}

//hello Shan la,
