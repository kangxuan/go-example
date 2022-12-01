package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 1-第一种声明方法，作为值类型
	var pers1 Person
	pers1.firstName = "kang"
	pers1.lastName = "xuan"
	upPerson(&pers1)
	fmt.Println(pers1)

	// 2-第二种声明方法，用new声明出来的为指针
	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	fmt.Println(pers2)
	(*pers2).lastName = "hhh" // 指针类型需要通过*取值
	fmt.Println(pers2)

	// 3-第三种声明方法，字面量声明
	pers3 := &Person{"kang", "bo"}
	pers3.lastName = "yi"
	upPerson(pers3)
	fmt.Println(pers3)
}

//{KANG XUAN}
//&{Chris Woodward}
//&{Chris hhh}
//&{KANG YI}
