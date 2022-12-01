package main

import (
	"fmt"
	"time"
)

type Address struct {
	Street           string // 街道
	HouseNumber      uint32 // 门牌号
	HouseNumberAddOn string
	POBox            string
	ZipCode          string
	City             string
	Country          string
}

type VCard struct {
	FirstName string
	LastName  string
	NickName  string
	BirtDate  time.Time
	Photo     string
	Addresses map[string]*Address // 可以有多个地址，采用指针会减少内存的消耗

}

func main() {
	addr1 := &Address{"Elfenstraat", 12, "", "", "2600", "Mechelen", "België"}
	addr2 := &Address{"Heideland", 28, "", "", "2640", "Mortsel", "België"}
	addrs := make(map[string]*Address)
	addrs["youth"] = addr1
	addrs["now"] = addr2

	birthdt := time.Date(1956, 1, 17, 15, 4, 5, 0, time.Local)
	photo := "MyDocuments/MyPhotos/photo1.jpg"
	vcard := &VCard{"Ivo", "Balbaert", "", birthdt, photo, addrs}
	fmt.Printf("Here is the full VCard: %v\n", vcard)
	fmt.Printf("My Addresses are:\n %v\n %v", addr1, addr2)
}

//Here is the full VCard: &{Ivo Balbaert  1956-01-17 15:04:05 +0800 CST MyDocuments/MyPhotos/photo1.jpg map[now:0xc00009a070 youth:0xc00009a000]}
//My Addresses are:
// &{Elfenstraat 12   2600 Mechelen België}
// &{Heideland 28   2640 Mortsel België}
