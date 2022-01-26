package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var opCount int
var mutex sync.Mutex

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
	mutex.Lock()
	{
		opCount++
	}
	mutex.Unlock()
	wg.Done()
}
