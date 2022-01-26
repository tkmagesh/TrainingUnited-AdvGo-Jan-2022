package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Counter struct {
	opCount int
	sync.Mutex
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.opCount++
	}
	c.Unlock()
}

var mutex sync.Mutex
var counter Counter

func main() {
	fmt.Println("main - entering")

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add(i, i)
	}
	wg.Wait()
	fmt.Println("opCount = ", counter.opCount)
	fmt.Println("main - exiting")
}

func add(x, y int) {
	//fmt.Println(x + y)
	counter.Increment()
	wg.Done()
}
