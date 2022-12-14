package main

import (
	"io"
	"log"
	"net/http"
)

type HandleFnc func(w http.ResponseWriter, request *http.Request)

const form2 = `<html><body><form action="#" method="post" name="bar">
		<input type="text" name="in"/>
		<input type="submit" value="Submit"/>
	</form></html></body>`

func main() {
	http.HandleFunc("/test1", logPanics(SimpleServer1))
	http.HandleFunc("/test2", logPanics(FormServer1))
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}

func SimpleServer1(w http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(w, "<h1>hello world</h1>")
	if err != nil {
		return
	}
}

func FormServer1(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		_, err := io.WriteString(w, form2)
		if err != nil {
			return
		}
	case "POST":
		_, err := io.WriteString(w, request.FormValue("in"))
		if err != nil {
			return
		}
	}
}

func logPanics(function HandleFnc) HandleFnc {
	return func(w http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		function(w, request)
	}
}
