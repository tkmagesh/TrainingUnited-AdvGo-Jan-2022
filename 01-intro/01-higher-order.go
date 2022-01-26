package main

import "fmt"

func main() {
	var fn func()

	fn = func() {
		fmt.Println("fn is invoked")
	}
	//fn()
	exec(fn)

	/*
		fmt.Println(add(100, 200))
		fmt.Println(subtract(100, 200))
	*/
	/*
		fmt.Println(logOperation(add, 100, 200))
		fmt.Println(logOperation(subtract, 100, 200))
	*/

	logAdd := getLogOperation(add)
	logSubtract := getLogOperation(subtract)
	fmt.Println(logAdd(100, 200))
	fmt.Println(logSubtract(100, 200))
}

/* func fn() {
	fmt.Println("fn is invoked")
} */

func exec(f func()) {
	f()
}

func logOperation(operation func(int, int) int, x, y int) int {
	fmt.Println("before invocation")
	result := operation(x, y)
	fmt.Println("after invocation")
	return result
}

func getLogOperation(operation func(int, int) int) func(x, y int) int {
	return func(x, y int) int {
		fmt.Println("before invocation")
		result := operation(x, y)
		fmt.Println("after invocation")
		return result
	}
}

func add(x, y int) int {
	result := x + y
	return result
}

func subtract(x, y int) int {
	result := x - y
	return result
}
