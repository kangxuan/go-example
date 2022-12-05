package main

import (
	"encoding/gob"
	"log"
	"os"
)

type Address2 struct {
	Type    string
	City    string
	Country string
}

type VCard2 struct {
	FirstName string
	LastName  string
	Addresses []Address2
	Remark    string
}

func main() {
	pa := Address2{"private", "Aartselaar", "Belgium"}
	wa := Address2{"work", "Boom", "Belgium"}
	vc := VCard2{"Jan", "Kersschot", []Address2{pa, wa}, "none"}

	//fmt.Printf("%v: \n", vc)
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println("关闭文件失败.")
			return
		}
	}(file)

	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
}
