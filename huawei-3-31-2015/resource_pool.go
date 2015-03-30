package main

import "math/rand"

func main() {
	sigs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		sigs[i] = make(chan int)
		go func(n int, sig chan int) {
			for {
				<-sig
				sig <- n + rand.Int()
			}
		}(i, sigs[i])
	}
	for i := 0; i < 100; i++ {
		sigs[i%10] <- 0
		println("got resource result ", <-sigs[i%10])
	}
}
