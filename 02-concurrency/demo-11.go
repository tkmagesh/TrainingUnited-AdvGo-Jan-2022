package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go genEvenNos(ch)
	for no := range ch {
		fmt.Println(no)
	}
}

func genEvenNos(ch chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i * 2
	}
	close(ch)
}
