package main

import "sync"

// startLoop starts a loop in a goroutine. the returned WaitGroup is done
// after the first loop iteration has started
func startLoop(n int) *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		first := true
		for {
			// do some work here
			if first {
				wg.Done()
				first = false
			}
		}
	}()
	return &wg
}
