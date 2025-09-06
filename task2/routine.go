package main

import (
	"fmt"
	"sync"
)

var wa sync.WaitGroup

func test1() {
	for i := 1; i < 10; i = i + 2 {
		fmt.Println("奇数:", i)
	}
	wa.Done() //goroutine 结束就登记-1
}

func main() {
	wa.Add(1) //登记一个goroutine
	go test1()
	for i := 2; i <= 10; i = i + 2 {
		fmt.Println("偶数:", i)
	}
	wa.Wait()
}
