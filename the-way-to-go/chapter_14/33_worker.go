package main

import (
	"fmt"
	"sync"
)

// Task 定义任务结构
type Task struct {
	TaskName string
}

// Pool 定义一个任务池
type Pool struct {
	Mu    sync.Mutex // 锁
	Tasks []*Task    // 任务列表
}

func Worker(pool *Pool) {
	for {
		pool.Mu.Lock()              // 加锁
		task := pool.Tasks[0]       // 取出一个任务
		pool.Tasks = pool.Tasks[1:] // 抛弃取出的任务
		pool.Mu.Unlock()            // 解锁
		exec(task)
		if len(pool.Tasks) == 0 {
			fmt.Println("任务已执行完毕")
			return
		}
	}
}

func exec(task *Task) {
	fmt.Printf("正在执行任务：%s\n", task.TaskName)
}

func main() {
	pools := Pool{
		Mu: sync.Mutex{},
		Tasks: []*Task{
			{
				TaskName: "task1",
			},
			{
				TaskName: "task2",
			},
			{
				TaskName: "task3",
			},
		},
	}

	Worker(&pools)
}
