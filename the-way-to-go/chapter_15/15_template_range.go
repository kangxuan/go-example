package main

import (
	"os"
	"text/template"
)

func main() {
	s := []int{1, 2, 3}
	t := template.New("test")
	t = template.Must(t.Parse("{{range .}}{{.}}\n{{end}}"))
	err := t.Execute(os.Stdout, s)
	if err != nil {
		return
	}
}

//1
//2
//3
