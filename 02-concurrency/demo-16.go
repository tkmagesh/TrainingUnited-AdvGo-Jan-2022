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
	elapsed := time.After(20 * time.Second)
	go func() {
		i := 0
		for {
			select {
			case <-elapsed:
				close(ch)
				return
			case ch <- i * 2:
				time.Sleep(500 * time.Millisecond)
				i++
			}
		}
	}()
	return ch
}
