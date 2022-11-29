package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%02d.%02d.%02d.%04d\n", t.Second(), t.Minute(), t.Hour(), t.Day(), t.Month(), t.Year())

	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())

	week = 60 * 60 * 24 * 7 * 1e9
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now)

	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
}

// 2022-11-28 17:49:19.85893 +0800 CST m=+0.000113248
//19.49.17.28.11.2022
//2022-11-28 09:49:19.859165 +0000 UTC
//2022-11-28 17:49:19.859169 +0800 CST m=+0.000352571
//2022-12-05 09:49:19.859165 +0000 UTC
//28 Nov 22 09:49 UTC
//Mon Nov 28 09:49:19 2022
//28 Nov 2022 09:49
//2022-11-28 09:49:19.859165 +0000 UTC => 20221128
