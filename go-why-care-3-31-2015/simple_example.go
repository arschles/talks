package main

import "fmt"

const numIters = 10

func main() {
	for i := 0; i < numIters; i++ {
		fmt.Printf("Hi Gophers at Huawei! (# %d)\n", i+1)
	}
}
