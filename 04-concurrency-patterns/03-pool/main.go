package main

import (
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"
)

type DBConnection struct {
	ID int
}

func (dbc *DBConnection) Close() error {
	fmt.Printf("Closing %d\n", dbc.ID)
	return nil
}

var IDCounter int

func DBConnectionFactory() (io.Closer, error) {
	IDCounter++
	fmt.Printf("DBConnectionFactory : Creating resource %d\n", IDCounter)
	return &DBConnection{ID: IDCounter}, nil
}

func main() {
	clientCount := 10
	p, err := pool.New(DBConnectionFactory /* factory */, 5 /* pool size */)
	wg := &sync.WaitGroup{}
	wg.Add(clientCount)
	for i := 0; i < clientCount; i++ {
		go func(client int) {
			doWork(client, p)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func doWork(client int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		fmt.Printf("Error %d: %s\n", client, err)
		return
	}
	defer p.Release(conn)
	fmt.Printf("Client %d: Acquired %d\n", client, conn.(*DBConnection).ID)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("Client %d Releasing %d\n", client, conn.(*DBConnection).ID)

}
