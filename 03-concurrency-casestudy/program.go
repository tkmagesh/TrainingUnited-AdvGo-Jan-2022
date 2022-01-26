package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

/*
Problem statement:
	Find the sum of all even numbers from both the files
	Find the sum of all odd numbers from both the files
	Write the sum of even numbers & sum of odd number into another file "result.txt"
*/
func main() {
	dataCh := make(chan int)
	fileWg := &sync.WaitGroup{}
	fileWg.Add(2)
	go Source("data1.dat", dataCh, fileWg)
	go Source("data2.dat", dataCh, fileWg)

	processWg := &sync.WaitGroup{}
	evenSumCh := make(chan int)
	oddSumCh := make(chan int)

	processWg.Add(3)
	evenCh, oddCh := Splitter(dataCh)
	go Sum(evenCh, evenSumCh, processWg)
	go Sum(oddCh, oddSumCh, processWg)
	go Merger(evenSumCh, oddSumCh, "result.txt", processWg)
	fileWg.Wait()
	close(dataCh)

	processWg.Wait()
	fmt.Println("Done")
}

func Source(filename string, dataCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if no, err := strconv.Atoi(txt); err == nil {
			dataCh <- no
		}
	}
}

func Splitter(dataCh chan int) (even, odd chan int) {
	even = make(chan int)
	odd = make(chan int)
	go func() {
		for no := range dataCh {
			if no%2 == 0 {
				even <- no
			} else {
				odd <- no
			}
		}
		close(even)
		close(odd)
	}()
	return
}

func Sum(noCh chan int, resultCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for no := range noCh {
		sum += no
	}
	resultCh <- sum
}

func Merger(evenSum, oddSumCh chan int, resultFile string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Create(resultFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for i := 0; i < 2; i++ {
		select {
		case evenSumVal := <-evenSum:
			writer.WriteString("Even total = " + strconv.Itoa(evenSumVal) + "\n")
		case oddSumVal := <-oddSumCh:
			writer.WriteString("Odd total = " + strconv.Itoa(oddSumVal) + "\n")
		}
	}
}
