package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers = 20
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	submitCh := make(chan workItem)
	startWorkers(ch, numWorkers)

	// submit work to each worker
	workItems := make([]workItem, numWorkers)
	for i := 0; i < numWorkers; i++ {
		wItem := workItem{
			reply: make(chan int),
			add:   rand.Int(),
			dur:   time.Duration(rand.Intn(10)) * time.Second,
		}
		submitCh <- wItem
		workItems[i] = wItem
	}

	// receive work from all the workers. results will receive as they are
	// completed by workers. use the sync.WaitGroup as a simple barrier to wait
	// until all receies to complete. for more on barriers, see ./barrier.go
	var wg sync.WaitGroup
	for _, wi := range workItems {
		wg.Add(1)
		go func(repl <-chan int) {
			defer wg.Done()
			fmt.Println(<-repl)
		}(wi.reply)
	}

	wg.Wait()

	// Note: we didn't build in a mechanism to shut down the wrk goroutines.
	// Hint: use the context package from ./ctx.go to do that!
}
