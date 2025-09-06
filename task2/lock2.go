package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// sync/atomic
var count1 int64
var w2 sync.WaitGroup

func test2() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&count1, 1)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go test2()
	}
	wg.Wait()
	fmt.Println("count=", count1)
}
