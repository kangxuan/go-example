package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Friend2 struct {
	FName string
}

type Person2 struct {
	UserName string
	Emails   []string
	Friends  []*Friend2
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	subStrs := strings.Split(s, "0")
	if len(subStrs) != 2 {
		return s
	}

	return subStrs[0] + " at " + subStrs[1]
}

func main() {
	f1 := Friend2{FName: "shan la 1"}
	f2 := Friend2{FName: "shan la 2"}

	t := template.New("filename example")
	// 定义模板中可以使用的函数
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`hello {{.UserName}}!
	{{range .Emails}}
		an emails {{.|emailDeal}}
	{{end}}
	{{with .Friends}}
	{{range .}}
		my friend name is {{.FName}}
	{{end}}
	{{end}}
	`)
	p := Person2{
		UserName: "shan la",
		Emails:   []string{"shanla0@qq.com", "shanla0@163.com"},
		Friends:  []*Friend2{&f1, &f2},
	}
	_ = t.Execute(os.Stdout, p)
}

//hello shan la!
//
//                an emails shanla at @qq.com
//
//                an emails shanla@163.com
//
//
//
//                my friend name is shan la 1
//
//                my friend name is shan la 2
