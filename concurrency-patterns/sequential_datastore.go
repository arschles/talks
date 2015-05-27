package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	ints := GetAll()
	fmt.Println(ints)
	fmt.Println("took", time.Now().Sub(start))
}

func datastoreGet() int {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return rand.Int()
}

func GetAll() []int {
	ints := make([]int, 10)
	for i := 0; i < 10; i++ {
		ints[i] = datastoreGet() // sleeps for <= 1sec, then returns a random int
	}
	return ints
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
