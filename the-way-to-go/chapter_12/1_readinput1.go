package main

import "fmt"

var (
	firstName, lastName, s, s1 string
	i                          int
	f                          float32
	input                      = "56.12 / 5212 / Go / s"
	format                     = "%f / %d / %s / %s"
	input1                     = "hh 2 4 6"
)

func main() {
	fmt.Println("Please enter your full name: ")
	//fmt.Scanln(&firstName, &lastName)
	fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s \n", firstName, lastName)

	fmt.Sscanf(input, format, &f, &i, &s, &s1)
	fmt.Println("From the string we read: ", f, i, s, s1)

	fmt.Sscanln(input1, &s, &s1)
	fmt.Println("From the input1 we read: ", s, s1)
}

//Please enter your full name:
//shan la
//Hi shanla
//From the string we read:  56.12 5212 Go s
//From the input1 we read:  hh 2
