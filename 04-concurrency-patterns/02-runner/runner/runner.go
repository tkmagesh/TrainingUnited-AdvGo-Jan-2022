package runner

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	t         time.Duration
	timeout   <-chan time.Time
	tasks     []func(int)
	complete  chan error
	interrupt chan os.Signal
}

var ErrTimeout = errors.New("timeout occurred")
var ErrInterrupt = errors.New("interrupt occurred")

func New(t time.Duration) *Runner {
	return &Runner{
		t:         t,
		tasks:     []func(int){},
		complete:  make(chan error),
		interrupt: make(chan os.Signal),
	}
}

func (r *Runner) Add(task func(int)) {
	r.tasks = append(r.tasks, task)
}

func (r *Runner) Start() error {
	r.timeout = time.After(r.t)
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	case <-r.interrupt:
		return ErrInterrupt
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		fmt.Println("Interrupt received... exiting")
		return true
	default:
		return false
	}
}
