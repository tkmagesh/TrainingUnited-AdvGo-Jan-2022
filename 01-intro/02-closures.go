package main

import "fmt"

func main() {
	increment, decrement := getSpinner()

	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4
	//counter = 1000
	fmt.Println(decrement()) //=> 3
	fmt.Println(decrement()) //=> 2
	fmt.Println(decrement()) //=> 1
	fmt.Println(decrement()) //=> 0
	fmt.Println(decrement()) //=> -1
}

func getSpinner() (func() int, func() int) {
	var counter int = 0

	increment := func() int {
		counter++
		return counter
	}

	decrement := func() int {
		counter--
		return counter
	}

	return increment, decrement
}
