package main

import (
	"fmt"
	"os"
	"regexp"
)

// IsIP 判断是否是IP
func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return false
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: regexp [string]")
		os.Exit(1)
	} else if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println("is number")
	} else {
		fmt.Println("is not number")
	}

	if IsIP("127.0.0.1") {
		fmt.Println("is IP")
	} else {
		fmt.Println("is not IP")
	}
}

//go run 7.go 1223sd
//is not number
//is not IP
