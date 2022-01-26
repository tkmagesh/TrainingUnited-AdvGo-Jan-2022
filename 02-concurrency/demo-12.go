package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genEvenNos()
	for no := range ch {
		fmt.Println(no)
	}
}

func genEvenNos() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 2
		}
		close(ch)
	}()
	return ch
}
