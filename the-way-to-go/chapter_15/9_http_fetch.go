package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://heljl.com/")
	checkError4(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError4(err)
	fmt.Printf("Got : %q", string(data))
}

func checkError4(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}

//2022/12/14 15:24:24 Get : Get "https://heljl.com/": dial tcp: lookup heljl.com: no such host
