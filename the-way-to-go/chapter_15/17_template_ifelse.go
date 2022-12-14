package main

import (
	"html/template"
	"os"
)

func main() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("Empty pipeline if demo: {{if ``}} Will not print. {{end}}\n"))
	err := tEmpty.Execute(os.Stdout, nil)
	if err != nil {
		return
	}

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("Non empty pipeline if demo: {{if `anything`}} will print. {{end}}\n"))
	err = tWithValue.Execute(os.Stdout, nil)
	if err != nil {
		return
	}

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if ``}} Print IF part. {{else}} Print ELSE  part. {{end}}\n"))
	err = tIfElse.Execute(os.Stdout, nil)
	if err != nil {
		return
	}
}

//Empty pipeline if demo:
//Non empty pipeline if demo:  will print.
//if-else demo:  Print ELSE  part.
