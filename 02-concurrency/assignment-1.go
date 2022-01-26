package main

import (
	"fmt"
	"sync"
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

var wg = &sync.WaitGroup{}

func main() {
	wg.Add(2)
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool)
	go print("Hello", ch1, ch2)
	go print("World", ch2, ch1)
	ch1 <- true
	wg.Wait()
}

func print(s string, in, out chan bool) {
	for i := 0; i < 5; i++ {
		<-in
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
		out <- true
	}
	wg.Done()
}
