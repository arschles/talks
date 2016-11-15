package main

import (
	"time"
)

func startActor(actorID string) (chan chan string, chan chan chan string) {
	strCh := make(chan chan string)
	fwdCh := make(chan chan chan string)
	go func() {
		for {
			select {
			case strRet := <-strCh:
				strRet <- actorID
			case fwd := <-fwdCh:
				// fwd is a chan chan string

			}

		}
	}()

}
