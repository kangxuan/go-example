package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

// 演示go-simplejson使用方法
func main() {
	js, err := simplejson.NewJson([]byte(`{
		"test": {
			"array": [1, "2", 3],
			"int": 10,
			"float": 5.150,
			"bignum": 9223372036854775807,
			"string": "simplejson",
			"bool": true
		}
	}`))
	if err != nil {
		return
	}

	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	s, _ := js.Get("test").Get("string").String()
	// 必须返回一直字符串，如果不符合则返回默认
	ms := js.Get("test").Get("bool").MustString("default_string")
	fmt.Println(arr)
	fmt.Println(i)
	fmt.Println(s)
	fmt.Println(ms)

}

//[1 2 3]
//10
//simplejson
//default_string
