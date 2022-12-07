package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	fmt.Println("In main()：")
	fmt.Println(startTime)
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	time.Sleep(10 * time.Second)
	fmt.Println("At the end of main()：")
	endTime := time.Now()
	fmt.Println(endTime)
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * time.Second) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * time.Second) // sleep for 5 seconds
	fmt.Println("End of shortWait()")
}

//In main()：
//2022-12-06 16:51:31.945567 +0800 CST m=+0.000087134
//About to sleep in main()
//Beginning shortWait()
//Beginning longWait()
//End of shortWait()
//End of longWait()
//At the end of main()：
//2022-12-06 16:51:41.947852 +0800 CST m=+10.002328696
