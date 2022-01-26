package main

import (
	"fmt"
	"time"
)

func main() {

	ch, done := genEvenNos()

	go func() {
		var input string
		fmt.Scanln(&input)
		done <- true
	}()

	for no := range ch {
		fmt.Println(no)
	}
}

func genEvenNos() (<-chan int, chan<- bool) {
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		i := 0
		for {
			select {
			case <-done:
				close(ch)
				return
			case ch <- i * 2:
				time.Sleep(500 * time.Millisecond)
				i++
			}
		}
	}()
	return ch, done
}
