package main

import (
	"fmt"
	"sync"
)

var mutx sync.Mutex
var count int
var w1 sync.WaitGroup

func test() {
	defer w1.Done()
	for i := 0; i < 1000; i++ {
		mutx.Lock()
		count++
		mutx.Unlock()
	}
}

func main() {
	for i := 0; i < 10; i++ {
		w1.Add(1)
		go test()
	}
	w1.Wait()
	fmt.Println("count=", count)
}
