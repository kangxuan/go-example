package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool   "This is bool type"
	field2 int    "This is int type"
	field3 string "This is string type"
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

func main() {
	tt := new(TagType)
	tt.field1 = false
	tt.field2 = 10
	tt.field3 = "呵呵"

	for i := 0; i < 3; i++ {
		refTag(*tt, i)
	}
}

//This is bool type
//This is int type
//This is string type
