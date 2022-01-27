package main

import (
	"fmt"
	"os"
	"runner-demo/runner"
	"time"
)

func main() {
	/*
		intialize a new runner with a timeout
		assign multiple tasks to the runner
		start the runner
		if all the tasks are completed with the given time, report "success"
		if the tasks are not completed with the given time, report "timeout"
		exit if the execution is interrupted by as os interrupt
	*/

	/*
		Sending signals
		https://www.cyberciti.biz/faq/unix-kill-command-examples/
	*/
	var input string
	fmt.Printf("Process %d started\n", os.Getpid())
	fmt.Println("Hit ENTER to continue...")
	fmt.Scanln(&input)

	timeout := 15 * time.Second
	r := runner.New(timeout)

	r.Add(createTask(3))
	r.Add(createTask(5))
	r.Add(createTask(8))

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			fmt.Println("timeout")
		case runner.ErrInterrupt:
			fmt.Println("interrupt")
		}
	} else {
		fmt.Println("success")
	}
}

func createTask(t int) func(int) {
	return func(id int) {
		fmt.Println("Processing task id ", id)
		time.Sleep(time.Duration(t) * time.Second)
	}
}
