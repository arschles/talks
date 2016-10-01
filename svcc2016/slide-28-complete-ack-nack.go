package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numWorkers = 100
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func wrk(ch <-chan chan int) {
	ret := <-ch
	time.Sleep(time.Second)
	ret <- rand.Int()
}

func main() {
	wrkCh := make(chan chan int)
	for i := 0; i < numWorkers; i++ {
		go wrk(wrkCh)
	}
	for i := 0; i < numWorkers; i++ {
		retCh := make(chan int)
		wrkCh <- retCh
		fmt.Println(<-retCh)
	}
}
