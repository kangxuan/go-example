package main

import (
	"fmt"
)

// Task1 定义任务结构
type Task1 struct {
	TaskName string
}

func sendWork(pending chan *Task1, taskName string) {
	pending <- &Task1{
		TaskName: taskName,
	}
}

func consumeWork(done chan *Task1) {
	for t := range done {
		fmt.Printf("已执行任务:%s\n", t.TaskName)
	}
	//for {
	//	select {
	//	case t := <-done:
	//		fmt.Printf("已执行任务:%s\n", t.TaskName)
	//	default:
	//		fmt.Println("没有任务了")
	//		return
	//	}
	//}
}

func exec1(task *Task1) {
	fmt.Printf("正在执行任务：%s\n", task.TaskName)
}

func Worker1(in, out chan *Task1) {
	//select {
	//case t := <-in:
	//	exec1(t)
	//default:
	//	out <- &Task1{
	//		TaskName: "结束了",
	//	}
	//}
	for {
		t := <-in
		exec1(t)
		out <- t
	}
}

func main() {
	pending, done := make(chan *Task1, 3), make(chan *Task1, 3)
	go sendWork(pending, "task1")
	go sendWork(pending, "task2")
	go sendWork(pending, "task3")
	for i := 0; i < 10; i++ { // start N goroutines to do work
		go Worker1(pending, done)
	}
	consumeWork(done)
}
