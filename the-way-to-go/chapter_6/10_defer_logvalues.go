package main

import "log"
import "io"

func main() {
	func1("Go")
}

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}
