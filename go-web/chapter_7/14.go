package main

import (
	"html/template"
	"os"
)

// 展示模板变量
func main() {
	t := template.New("")
	t, _ = t.Parse(`{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
{{with $x := "output"}}{{printf "%q"  $x}}{{end}}
{{with $x := "output"}}{{$x | printf "%q"}}{{end}}
	`)
	t.Execute(os.Stdout, nil)
}

//&#34;output&#34;
//&#34;output&#34;
//&#34;output&#34;
