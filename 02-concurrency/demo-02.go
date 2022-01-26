package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main started")
	wg.Add(1)
	go f1() //scheduling to execute
	f2()
	wg.Wait()
	fmt.Println("main completed")
}

func f1() {
	defer wg.Done()
	fmt.Println("f1 invocation started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 invocation completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
