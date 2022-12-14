package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入网址")

	url, err := inputReader.ReadString('\n')
	url = strings.TrimSuffix(url, "\n")
	url = strings.TrimSpace(url)
	checkError(err)

	res, err := http.Get(url)
	checkError(err)

	data, err := ioutil.ReadAll(res.Body)
	checkError(err)

	fmt.Printf("Got : %q", string(data))

}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}

//请输入网址
//https://baidu.com
//Got : "<!DOCTYPE html><!--STATUS OK--><html><head><meta http-equiv=\"Content-Type\" content=\"text/html;charset=utf-8\"><meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge,chrome=1\"><meta content=\"always\" name=\"referrer\"><meta name=\"theme-color\" content=\"#ffffff\"><meta name=\"description\" content=\"全球领先的中文搜索引擎、致力...
