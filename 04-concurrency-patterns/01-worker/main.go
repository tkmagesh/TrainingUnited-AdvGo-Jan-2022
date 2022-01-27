package main

import (
	"fmt"
	"time"
	"worker-demo/worker"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
	"Magesh",
	"Ganesh",
	"Ramesh",
	"Rajesh",
	"Suresh",
}

type NamePrinter struct {
	name  string
	delay time.Duration
}

/* worker.Work interface implementation */
func (np *NamePrinter) Task() {
	fmt.Println("Task execution commenced")
	time.Sleep(np.delay)
	fmt.Println("Name Printer - Name : ", np.name)
}

func main() {
	wkr := worker.New(5 /* batch size */)
	for idx := 0; idx < 2; idx++ {
		for index, name := range names {
			np := &NamePrinter{
				name:  name,
				delay: time.Duration(index*100) * time.Millisecond,
			}
			wkr.Run(np)
		}
	}
	fmt.Println("All tasks are assigned")
	wkr.Shutdown() /* wait for all the tasks to complete and shutdown */
}
