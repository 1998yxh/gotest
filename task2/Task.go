package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

type Scheduler struct {
	tasks []Task
}

func (s *Scheduler) AddTask(t Task) {
	s.tasks = append(s.tasks, t)
}

func (s *Scheduler) Run() {
	var wg sync.WaitGroup
	for i, task := range s.tasks {
		wg.Add(1)
		go func(idx int, t Task) {
			defer wg.Done()
			start := time.Now()
			t()
			elapsed := time.Since(start)
			fmt.Printf("任务%d执行时间: %v\n", idx, elapsed)
		}(i, task)
	}
	wg.Wait()
}

func main() {
	s := &Scheduler{}
	s.AddTask(func() { time.Sleep(500 * time.Millisecond) })
	s.AddTask(func() { time.Sleep(300 * time.Millisecond) })
	s.AddTask(func() { time.Sleep(700 * time.Millisecond) })
	s.Run()
}
