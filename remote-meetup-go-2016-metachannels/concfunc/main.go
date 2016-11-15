package main

import (
	"log"
)

func main() {
	oneCh := make(chan arg)
	twoCh := make(chan arg)
	go adder(1, oneCh)
	go adder(2, twoCh)
}
