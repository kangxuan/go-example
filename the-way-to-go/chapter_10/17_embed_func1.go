package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String1() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

func main() {
	c := &Customer{"康宣", &Log{"第一条日志"}}
	c.log.Add("第二条日志")
	fmt.Println(c.Log())
}

//&{第一条日志
//第二条日志}
