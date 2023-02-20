package main

import "github.com/cihub/seelog"

func main() {
	defer seelog.Flush() // 程序结束前刷入

	seelog.Info("Hello from Seelog")
}
