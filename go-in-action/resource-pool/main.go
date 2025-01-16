package main

import (
	"fmt"
	"io"
	"math/rand"
	"pool/pool"
	"sync"
	"sync/atomic"
	"time"
)

var counter uint64
var wg sync.WaitGroup

const (
	poolSize = 5
	jobs     = 40
)

type connection struct {
	id uint64
}

func (c *connection) Close() error {
	fmt.Println("Closing connection", c.id)
	return nil
}

func main() {
	p, _ := pool.New(createConnection, poolSize)
	wg.Add(jobs)

	for i := 0; i < jobs; i++ {
		go func(i int) {
			if err := doWork(i, p); err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func createConnection() (io.Closer, error) {
	atomic.AddUint64(&counter, 1)

	return &connection{
		counter,
	}, nil
}

func doWork(query int, p *pool.Pool) error {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	conn, err := p.Acquire()

	if err != nil {
		return err
	}

	defer p.Release(conn)

	fmt.Println("Working with connection", conn.(*connection).id, "on query", query)

	return nil
}
