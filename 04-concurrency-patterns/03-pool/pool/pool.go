package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

type Pool struct {
	factory   func() (io.Closer, error)
	resources chan io.Closer
	mutex     *sync.Mutex
	closed    bool
}

var ErrInvalidPoolSize = errors.New("invalid pool size")
var ErrPoolClosed = errors.New("pool closed")

func New(factory func() (io.Closer, error), poolSize int) (*Pool, error) {
	if poolSize <= 0 {
		return nil, ErrInvalidPoolSize
	}
	return &Pool{
		factory:   factory,
		resources: make(chan io.Closer, poolSize),
		mutex:     &sync.Mutex{},
		closed:    false,
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case resource, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		fmt.Println("Acquiring from the pool")
		return resource, nil
	default:
		fmt.Println("Creating a new resource from factory")
		return p.factory()
	}
}

func (p *Pool) Release(resource io.Closer) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.closed {
		resource.Close()
		return ErrPoolClosed
	}

	select {
	case p.resources <- resource:
		fmt.Println("Releasing into the pool")
		return nil
	default:
		fmt.Println("Releasing the resource: closing & discarding")
		return resource.Close()
	}
}

func (p *Pool) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.resources)
	for resource := range p.resources {
		resource.Close()
	}
}
