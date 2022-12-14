package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprint(w, "Hello!")
	if err != nil {
		return
	}
}

func main() {
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		return
	}
}
