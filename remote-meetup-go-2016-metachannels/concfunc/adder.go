package main

import (
	"time"
)

type arg struct {
	i   int
	ret chan int
}

// adder adds offset to every argument sent to it over argsCh, and responds with the result over
// the ret channel passed as part of the arg. Does a bit more "processing" and then closes ret
//
// The response is done asynchronously, so if the 'caller'
// (code that sent on argsCh) doesn't receive on the ret channel, this function won't block
func adder(offset int, argsCh <-chan arg) {
	for {
		select {
		case a := <-argsCh:
			go func(a arg) {
				a.ret <- a.i + offset
				time.Sleep(1 * time.Second)
				close(a.ret)
			}(a)
		}
	}
}
