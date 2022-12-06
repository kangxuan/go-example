package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func main() {
	var b []byte

	hash := sha1.New()
	_, err := io.WriteString(hash, "test")
	if err != nil {
		return
	}

	fmt.Printf("Result: %x\n", hash.Sum(b))
	fmt.Printf("Result: %d\n", hash.Sum(b))

	hash.Reset()

	data := []byte("We shall overcome!")
	n, err1 := hash.Write(data)
	if n != len(data) || err1 != nil {
		log.Printf("Hash write error: %v / %v", n, err1)
	}
	checksum := hash.Sum(b)
	fmt.Printf("Result: %x\n", checksum)

}

//Result: a94a8fe5ccb19ba61c4c0873d391e987982fbbd3
//Result: [169 74 143 229 204 177 155 166 28 76 8 115 211 145 233 135 152 47 187 211]
//Result: e2222bfc59850bbb00a722e764a555603bb59b2a
