package main

import (
	"io"
	"net/http"
)

const form = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="submit" value="submit"/>
		</form>
	</body></html>
`

func SimpleServer(w http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(w, "<h1>hello, world</h1>")
	if err != nil {
		return
	}
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		_, err := io.WriteString(w, form)
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

func main() {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}
