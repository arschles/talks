package main

import (
	"fmt"
	"time"
)

func call(every, until time.Duration, ch chan<- int, fn func(int)) {
	defer close(ch)
	if until < every {
		return
	}
	timer := time.NewTimer(until)
	ticker := time.NewTicker(every)
	defer timer.Stop()
	defer ticker.Stop()
	i := 0
	for {
		select {
		case <-timer.C:
			return
		case <-ticker.C:
			fn(i)
			ch <- i
			i++
		}
	}
}

func main() {
	every := time.Second
	until := 10 * time.Second
	ch := make(chan int)
	go call(every, until, ch, func(i int) {
		fmt.Println("call number", i)
	})
	for i := range ch {
		fmt.Println("called", i)
	}
}
