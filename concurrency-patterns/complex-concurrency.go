package main

import (
	"sync"

	"github.com/iron-io/go/ironmq/backend"
)

type pushQueueMap struct {
	sync.RWMutex
	m     map[string]*pushQueueInfo
	C     chan *dequeueSignal
	OC    chan *pushMapOp
	rc    chan bool // Channel to run carrier
	clock backend.Time
}

var lookup = make(map[string]int)
var lookupM = sync.RWMutex

func loop() {

}

func main() {

}
