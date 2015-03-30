package main

import (
	"fmt"
	"time"
)

const numIters = 10

func printHello(iterNum int) {
	fmt.Printf("Hi Gophers at Huawei! (# %d)\n", iterNum+1)
}

func main() {
	for i := 0; i < numIters; i++ {
		go printHello(i)
	}

	time.Sleep(1 * time.Second)
}
