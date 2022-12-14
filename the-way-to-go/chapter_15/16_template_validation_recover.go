package main

import (
	"fmt"
	"html/template"
	"log"
)

func main() {
	tOk := template.New("ok")
	tErr := template.New("error_template")
	defer func() {
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()

	template.Must(tOk.Parse("/* and a comment */ some static text: {{ .Name }}"))
	fmt.Println("The first one parsed OK.")
	fmt.Println("The next one ought to fail.")
	template.Must(tErr.Parse(" some static text {{ .Name }"))
}

//The first one parsed OK.
//The next one ought to fail.
//2022/12/14 16:11:39 run time panic: template: error_template:1: unexpected "}" in operand
