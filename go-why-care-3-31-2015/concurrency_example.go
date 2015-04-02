package main

import (
	"fmt"
	"sync"
)

const numIters = 10

func printHello(iterNum int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Hi Gophers! (# %d)\n", iterNum+1)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < numIters; i++ {
		wg.Add(1)
		go printHello(i, &wg)
	}
	wg.Wait()
}
