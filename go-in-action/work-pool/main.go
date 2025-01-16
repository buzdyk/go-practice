package main

import (
	"fmt"
	"math/rand"
	"pool/pool"
	"sync"
	"sync/atomic"
	"time"
)

var (
	counter int64
	wg      sync.WaitGroup
)

type task struct {
	id int64
}

func (w *task) Task() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println("working task", w.id)
}

func main() {
	p := pool.New(10)

	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func() {
			p.Run(createTask())
			wg.Done()
		}()
	}
	wg.Wait()
	p.Shutdown()
}

func createTask() *task {
	atomic.AddInt64(&counter, 1)
	return &task{
		counter,
	}
}
