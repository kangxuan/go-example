package main

import (
	"fmt"
)

// Task1 定义任务结构
type Task1 struct {
	TaskName string
}

// sendWork 发送任务到队列
func sendWork(pending chan *Task1, taskName string) {
	pending <- &Task1{
		TaskName: taskName,
	}
}

// 监听任务是否结束
func consumeWork(done chan *Task1) {
	select {
	case <-done:
		fmt.Println("没有任务了")
		return
	}
}

// 执行任务
func exec1(task *Task1) {
	fmt.Printf("正在执行任务：%s\n", task.TaskName)
}

// Worker1 Worker..
func Worker1(in, out chan *Task1) {
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
