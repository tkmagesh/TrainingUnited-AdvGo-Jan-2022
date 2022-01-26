package main

import "fmt"

type Employee struct {
	name string
	age  int
	id   int
}

type FulltimeEmployee struct {
	Employee
	benefits string
}

func main() {
	ft := FulltimeEmployee{Employee: Employee{name: "John", age: 30, id: 1}, benefits: "health insurance"}
	fmt.Printf("%#v\n", ft)
	//fmt.Println(ft.Employee.name)
	fmt.Println(ft.name)
}
