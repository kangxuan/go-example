package main

import (
	"fmt"
)

// 发动机接口
type Engine interface {
	Start() // 启动
	Stop()  // 停止
}

// 汽车结构
type Car struct {
	wheelCount int // 轮子数量
	Engine         // 引擎
}

// 定义Car的方法-获取轮子数量
func (c Car) NumberOfWheels() int {
	return c.wheelCount
}

// 定义奔驰结构
type Mercedes struct {
	Car
}

// 定义奔驰的方法
func (m Mercedes) sayHiToMerkel() {
	fmt.Println("Hi Angela!")
}

// 实现Car的start方法
func (c Car) Start() {
	fmt.Println("Car is started")
}

// 实现Car的stop方法
func (c Car) Stop() {
	fmt.Println("Car is Stopped")
}

// 增加Car的工作方法
func (c Car) GoToWorkIn() {
	c.Start()
	c.Stop()
}

func main() {
	m := &Mercedes{Car{4, nil}}
	fmt.Println("Mercedes wheelCount is:", m.NumberOfWheels())
	m.sayHiToMerkel()
	m.GoToWorkIn()
}

//Mercedes wheelCount is: 4
//Hi Angela!
//Car is started
//Car is Stopped
