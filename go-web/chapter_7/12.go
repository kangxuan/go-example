package main

import (
	"html/template"
	"os"
)

func main() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}}不会输出{{end}}\n"))
	_ = tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test1")
	template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}}输出内容{{end}}\n"))
	_ = tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test2")
	template.Must(tIfElse.Parse("if else demo: {{if `anything`}}if部分{{else}}else部分{{end}}\n"))
	_ = tIfElse.Execute(os.Stdout, nil)
}

//空 pipeline if demo:
//不为空的 pipeline if demo: 输出内容
//if else demo: if部分
