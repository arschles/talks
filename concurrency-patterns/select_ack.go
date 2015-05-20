package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doWork(calc <-chan chan<- int, cur <-chan chan<- int, history <-chan chan<- []int, stats chan<- int) {
	var all []int
	i := 0
	for {
		select {
		case ch := <-calc:
			i := rand.Int()
			ch <- i
			all = append(all, i)
		case ch := <-cur:
			ch <- i
		case ch := <-history:
			ch <- all
		}
		i++
		stats <- i
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	calc := make(chan chan<- int)
	cur := make(chan chan<- int)
	hist := make(chan chan<- []int)
	stats := make(chan int)
	go doWork(calc, cur, hist, stats)

	calcRes := make(chan int)
	go func() {
		calc <- calcRes
	}()
	fmt.Println(<-calcRes)
	fmt.Println(<-stats)
}
