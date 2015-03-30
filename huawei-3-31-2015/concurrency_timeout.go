package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "goroutine 1"
	}()
	select {
	case str := <-ch:
		fmt.Println("got string ", str)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out on goroutine1")
	}

}
