package main

type arg struct {
	i   int
	ret chan int
}

// adder adds offset to every argument sent to it over argsCh, and responds with the result over
// the ret channel passed as part of the arg.
//
// The response is done asynchronously, so if the 'caller'
// (code that sent on argsCh) doesn't receive on the ret channel, this function won't block
func adder(offset int, argsCh <-chan arg) {
	for {
		select {
		case a := <-argsCh:
			go func(a arg) {
				a.ret <- a.i + offset
			}(a)
		}
	}
}
