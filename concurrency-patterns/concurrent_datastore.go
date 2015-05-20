package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	start := time.Now()
	ints := GetAll()

	fmt.Println(ints)
	fmt.Println("took", time.Now().Sub(start))
}

func datastoreGet(ch chan<- int, wg *sync.WaitGroup) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	ch <- rand.Int()
}

func getWGAndChan(n int) (*sync.WaitGroup, chan int) {
	var wg sync.WaitGroup
	wg.Add(n)
	ch := make(chan int)
	return &wg, ch
}

func GetAll() []int {
	wg, ch := getWGAndChan(10) // HL
	for i := 0; i < 10; i++ {
		c := make(chan int)    // HL
		go datastoreGet(c, wg) // HL
		go func() {            // HL
			defer wg.Done() // HL
			ch <- <-c       // HL
		}() // HL
	}
	go func() { // HL
		wg.Wait() // HL
		close(ch) // HL
	}() // HL
	ints := make([]int, 10)
	i := 0
	for res := range ch {
		ints[i] = res
		i++
	}
	return ints
}
