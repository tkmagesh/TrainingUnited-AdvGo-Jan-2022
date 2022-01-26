package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	fmt.Println(<-ch)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
}
