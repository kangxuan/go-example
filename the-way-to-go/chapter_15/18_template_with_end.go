package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("test")
	t, _ = t.Parse("{{with `hello`}}{{.}}{{end}}!\n")
	err := t.Execute(os.Stdout, nil)
	if err != nil {
		return
	}

	t, _ = t.Parse("{{with `hello`}}{{.}} {{with `Mary`}}{{.}}{{end}}{{end}}!\n")
	err = t.Execute(os.Stdout, nil)
	if err != nil {
		return
	}
}

//hello!
//hello Mary!
