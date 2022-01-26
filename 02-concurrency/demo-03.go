package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//var wg sync.WaitGroup
	wg := &sync.WaitGroup{}
	fmt.Println("main started")
	wg.Add(1)
	//go f1(&wg) //scheduling to execute
	go f1(wg)
	f2()
	wg.Wait()
	fmt.Println("main completed")
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("f1 invocation started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 invocation completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
