package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sendNum(ch chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("打印整数:", i)
	}
	close(ch)

}

func getNum(ch chan int) {
	for v := range ch {
		fmt.Println("接受整数:", v)
	}
	wg.Done()
}

func main() {
	ch := make(chan int, 5)
	wg.Add(1)
	go getNum(ch)
	wg.Add(1)
	go sendNum(ch)
	wg.Wait()
}
