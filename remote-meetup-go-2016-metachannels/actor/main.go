package main

import (
	"log"
	"time"
)

func main() {
	fwdCh := startActor()

	sleepRes := make(chan time.Duration)
	sleepCh <- sleepRes
	log.Printf("slept for %s", <-sleepRes)

}
