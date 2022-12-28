package main

import (
	"html/template"
	"os"
)

type Friend struct {
	FName string
}

type Person1 struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{FName: "first"}
	f2 := Friend{FName: "second"}
	t := template.New("template 2")
	t, _ = t.Parse(`hello {{.UserName}}!
			{{range .Emails}}
				an email {{.}}
			{{end}}
			{{with .Friends}}
			{{range .}}
				my friend name is {{.FName}}
			{{end}}
			{{end}}
	`)
	p := Person1{
		UserName: "Shan la",
		Emails:   []string{"Shanla@qq.com", "Shanla@163.com"},
		Friends:  []*Friend{&f1, &f2},
	}
	_ = t.Execute(os.Stdout, p)
}

//hello Shan la!
//
//                                an email Shanla@qq.com
//
//                                an email Shanla@163.com
//
//
//
//                                my friend name is first
//
//                                my friend name is second
//
//
//
