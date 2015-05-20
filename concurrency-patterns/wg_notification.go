package main

import "sync"

// startLoop starts a loop in a goroutine. the returned WaitGroup is done
// after the first loop iteration has started
func startLoop(n int) *sync.WaitGroup {
	var wg sync.WaitGroup
	go func() {
		first := true
		for {
			if first {
				wg.Done()
				first = false
			}
			// do some work here
		}
	}()
	return &wg
}
