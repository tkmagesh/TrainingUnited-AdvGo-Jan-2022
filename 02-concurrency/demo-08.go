package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go add(100, 200, ch, wg)
	fmt.Println(<-ch)
	wg.Wait()
}

func add(x, y int, ch chan int, wg *sync.WaitGroup) {
	result := x + y
	ch <- result
	wg.Done()
}
