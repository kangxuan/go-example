package main

import "fmt"

var icons map[string]string

func loadIcons() {
	icons = map[string]string{
		"left":  "left.png",
		"up":    "up.png",
		"right": "right.png",
		"down":  "down.png",
	}
}

// Icon 被多个goroutine调用时不是并发安全的
func Icon(name string) {
	if icons == nil {
		loadIcons()
	}
	fmt.Println(icons[name])
}

func main() {
	for i := 0; i < 1000; i++ {
		go Icon("left")
	}
}
