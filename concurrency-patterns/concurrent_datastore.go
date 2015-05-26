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

func datastoreGet(ch chan<- int) {
	defer close(ch)
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
	wg, ch := getWGAndChan(10) // get a waitgroup that has 10 added to it, and a chan int // HL
	for i := 0; i < 10; i++ {
		c := make(chan int) // HL
		go datastoreGet(c)  // sends an int on c then closes after sleeping <= 1 sec // HL
		go func() {         // HL
			defer wg.Done() // mark this iteration done after receiving on c // HL
			ch <- <-c       // enhancement: range of c if >1 results // HL
		}() // HL
	}
	go func() { // HL
		wg.Wait() // wait for all datastoreGets to finish // HL
		close(ch) // then close the main channel // HL
	}() // HL
	ints := make([]int, 10)
	i := 0
	for res := range ch { // stream all results from each datastoreGet into the slice
		ints[i] = res // GetAll can be a generator if you're willing to change API.
		i++           // that lets you push results back to the caller.
	}
	return ints
}
