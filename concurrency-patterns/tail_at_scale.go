package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/net/context"
)

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	for i := 0; i < 10; i++ {
		// get sleeps for a random duration <= 100ms,
		// then sends a random int on ch. stops if ctx.Done() receives.
		go get(ctx, ch)
	}
	select {
	case i := <-ch:
		fmt.Printf("got result %d\n", i)
	case <-ctx.Done():
		fmt.Println("got no result")
	}
}

func get(ctx context.Context, ch chan<- int) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	select {
	case ch <- rand.Int():
	case <-ctx.Done():
		return
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
