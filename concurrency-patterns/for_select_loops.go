package main

import (
	"sync"

	"golang.org/x/net/context"
)

func RunLoops(ctx context.Context, n int) (*sync.WaitGroup, *sync.WaitGroup) {
  var startWG sync.WaitGroup
  var stopWG sync.WaitGroup
  startWG.Add(n)
  stopWG.Add(n)
	for i := 0; i < n; i++ {
    go func(grNum int) {
      for {
        select
      }
    }(i)
	}
}
