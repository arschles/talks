package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/net/context"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	ch := make(chan string)
	// simulates polling a queue. sends dequeued results on ch.
	// stops and closes ch when ctx.Done() receives
	go poll(ctx, ch)
	for polled := range ch {
		fmt.Println(polled)
	}
}

// poll simulates polling a queue. sends dequeued results on ch.
// stops and closes ch when ctx.Done() sends
func poll(ctx context.Context, ch chan<- string) {
	defer close(ch)
	for {
		time.Sleep(5 * time.Millisecond)
		select {
		case <-ctx.Done():
			return
		case ch <- "element " + strconv.Itoa(rand.Int()):
		}
	}
}
