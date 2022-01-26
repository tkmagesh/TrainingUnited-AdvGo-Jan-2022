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
	elapsed := after(20 * time.Second)
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

func after(d time.Duration) <-chan bool {
	doneCh := make(chan bool)
	go func() {
		time.Sleep(d)
		doneCh <- true
	}()
	return doneCh
}
