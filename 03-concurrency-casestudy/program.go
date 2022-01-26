package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
Problem statement:
	Find the sum of all even numbers from both the files
	Find the sum of all odd numbers from both the files
	Write the sum of even numbers & sum of odd number into another file "result.txt"
*/
func main() {
	var evenSum, oddSum int
	file, err := os.Open("data2.dat")
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if no, err := strconv.Atoi(txt); err == nil {
			if no%2 == 0 {
				evenSum += no
			} else {
				oddSum += no
			}
		}
	}
	fmt.Println(evenSum, oddSum)
}
