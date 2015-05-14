package main

import "fmt"

const numIters = 10

func main() {
	for i := 0; i < numIters; i++ {
		fmt.Printf("Hi Gophers! (# %d)\n", i+1)
	}
}
