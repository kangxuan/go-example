package main

import "fmt"

type specialString string

var whatIsThis specialString = "hello"

func TypeSwitch() {
	TestFunc := func(any interface{}) {
		switch v := any.(type) {
		case int:
			fmt.Printf("any %v is an int type", v)
		case float32:
			fmt.Printf("any %v is a float32 type", v)
		case string:
			fmt.Printf("any %v is a string type", v)
		case specialString:
			fmt.Printf("any %v is a special String!", v)
		default:
			fmt.Println("unknown type!")
		}
	}

	TestFunc(whatIsThis)
}

func main() {
	TypeSwitch()
}

//any hello is a special String!
