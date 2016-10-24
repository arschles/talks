package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numWorkers = 20
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type workItem struct {
	reply chan int
	add   int
	dur   time.Duration
}

func wrk(ch <-chan workItem) {
	item := <-ch                        // get some work to do
	time.Sleep(item.dur)                // do the "work"
	item.reply <- rand.Int() + item.add // return the result of the "work"
}

func main() {
	// start up workers, each waiting for a submission
	submitCh := make(chan workItem)
	for i := 0; i < numWorkers; i++ {
		go wrk(submitCh)
	}

	for i := 0; i < numWorkers; i++ {
		wi := workItem{
			reply: make(chan int),
			add:   rand.Int(),
			dur:   time.Duration(rand.Intn(150)) * time.Millisecond,
		}
		submitCh <- wi
		fmt.Println(<-wi.reply)
	}

	// Note: we don't have a mechanism to shut down the wrk goroutines in a
	// clean way. Use the context example in ./ctx.go to do that!
}
