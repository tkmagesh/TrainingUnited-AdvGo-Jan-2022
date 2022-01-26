package main

import (
	"fmt"
	"time"
)

/*
Expected result:
Hello
World
Hello
World
Hello
World
Hello
World
Hello
World

Important :
	The loop has to in the respective function

*/

func main() {
	printHello()
	printWorld()
}

func printHello() {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Hello")
	}
}

func printWorld() {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("World")
	}
}
