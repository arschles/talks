package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// does some "work" and sends the result on ch
func process(dur time.Duration, ch chan<- int) {
	time.Sleep(dur)
	ch <- rand.Int()
}

func main() {
	// this is the barrier
	wg := new(sync.WaitGroup)
	// this is the channel over which we stream results from the worker goroutines
	mainCh := make(chan int)

	for i := 0; i <= 20; i++ {
		wg.Add(1)
		ch := make(chan int)
		workDur := time.Duration(rand.Intn(20)) * time.Second
		go process(workDur, ch)

		go func() {
			// indicate that the worker goroutine is done after it processes and sends its result
			// to the main channel
			defer wg.Done()
			it := <-ch
			mainCh <- it
		}()
	}

	// close the main channel after the barrier is released
	go func() {
		wg.Wait()
		close(mainCh)
	}()

	// stream results from each goroutine as they process
	for i := range mainCh {
		log.Println(i)
	}
}
