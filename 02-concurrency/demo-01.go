package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	go f1() //scheduling to execute
	f2()
	/*
		var input string
		fmt.Scanln(&input)
	*/
	time.Sleep(10 * time.Millisecond) //never assume
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 invocation started")
	time.Sleep(5 * time.Millisecond)
	fmt.Println("f1 invocation completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
