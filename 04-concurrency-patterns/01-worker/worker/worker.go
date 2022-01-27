package worker

import (
	"fmt"
	"sync"
)

type Work interface {
	Task()
}

type Worker struct {
	work chan Work
	wg   *sync.WaitGroup
}

func New(batchSize int) *Worker {
	worker := &Worker{
		work: make(chan Work), /* why NOT a buffered a channel? */
		wg:   new(sync.WaitGroup),
	}
	worker.wg.Add(batchSize)
	for i := 0; i < batchSize; i++ {
		go func() {
			for w := range worker.work {
				w.Task()
			}
			worker.wg.Done()
		}()
	}
	return worker
}

func (worker *Worker) Run(w Work) {
	worker.work <- w
}

func (worker *Worker) Shutdown() {
	close(worker.work)
	worker.wg.Wait()
	fmt.Println("worker shutdown complete")
}
