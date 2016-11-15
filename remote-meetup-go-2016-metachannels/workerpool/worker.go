package main

import (
	"math/rand"
	"time"
)

type workItem struct {
	reply chan int
	add   int
	dur   time.Duration
}

// this is the worker func. N of these will be launched in the worker pool
func wrk(ch <-chan workItem) {
	item := <-ch                        // get some work to do
	time.Sleep(item.dur)                // do the work
	item.reply <- rand.Int() + item.add // return the result of the work
}
