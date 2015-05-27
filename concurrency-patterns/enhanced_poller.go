package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/context"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func makeThings(i int) (chan string, *sync.WaitGroup) {
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(i)
	return ch, &wg
}

func printCh(ch <-chan string) {
	for s := range ch {
		fmt.Println(s)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	mainCh, wg := makeThings(10) // make a chan string and a wg that has 10 added to it
	for i := 0; i < 10; i++ {
		start, _, ch := poll(ctx)
		start.Wait()
		go func() {
			defer wg.Done()
			for str := range ch {
				mainCh <- str
			}
		}()
	}
	go func() {
		wg.Wait()
		close(mainCh)
	}()
	printCh(mainCh) // loops on mainCh until it's closed
}

func poll(ctx context.Context) (*sync.WaitGroup, *sync.WaitGroup, <-chan string) {
	var start, end sync.WaitGroup // start & end notifications to multiple parties
	start.Add(1)
	end.Add(1)
	ch := make(chan string)
	go func() {
		defer close(ch)
		defer end.Done()
		start.Done()
		for {
			time.Sleep(5 * time.Millisecond)
			select {
			case <-ctx.Done():
				return
			case ch <- "element " + strconv.Itoa(rand.Int()):
			}
		}
	}()
	return &start, &end, ch
}
