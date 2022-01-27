package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

type Pool struct {
	factory          func() (io.Closer, error)
	resources        chan io.Closer
	mutex            *sync.Mutex
	size             int
	closed           bool
	resourcesCreated int
}

var ErrInvalidPoolSize = errors.New("invalid pool size")
var ErrPoolClosed = errors.New("pool closed")

func New(factory func() (io.Closer, error), poolSize int) (*Pool, error) {
	if poolSize <= 0 {
		return nil, ErrInvalidPoolSize
	}
	return &Pool{
		factory:          factory,
		resources:        make(chan io.Closer, poolSize),
		mutex:            &sync.Mutex{},
		size:             poolSize,
		closed:           false,
		resourcesCreated: 0,
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	/* To limit the # of resources created to the pool size */
	p.mutex.Lock()
	{
		if p.resourcesCreated < p.size {
			fmt.Println("Acquire : from factory")
			resource, err := p.factory()
			if err != nil {
				p.mutex.Unlock()
				return nil, err
			}
			p.resourcesCreated++
			p.resources <- resource
		}
	}
	p.mutex.Unlock()
	resource, ok := <-p.resources
	if !ok {
		return nil, ErrPoolClosed
	}
	return resource, nil

	/* Creating unlimited number of resources */
	/* select {
	case resource, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		fmt.Println("Acquiring from the pool")
		return resource, nil
	default:
		fmt.Println("Creating a new resource from factory")
		return p.factory()
	} */
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
