package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("test")
	t = template.Must(t.Parse("{{with $x := `hello`}}{{printf `%s %s` $x `Mary`}}{{end}}!\n"))
	err := t.Execute(os.Stdout, nil)
	if err != nil {
		return
	}
}

//hello Mary!
