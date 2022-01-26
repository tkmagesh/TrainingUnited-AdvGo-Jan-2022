package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

//Communicate by sharing memory (DON'T)
var opCount int32

func main() {
	fmt.Println("main - entering")

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add(i, i)
	}
	wg.Wait()
	fmt.Println("opCount = ", opCount)
	fmt.Println("main - exiting")
}

func add(x, y int) {
	//fmt.Println(x + y)
	atomic.AddInt32(&opCount, 1)
	wg.Done()
}
