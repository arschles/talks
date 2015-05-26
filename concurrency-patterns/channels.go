package main

import "sync"

// WatchChanges will watch the state of the given request. ch will send after
// each request state change and will be closed after the request is removed from
// the request state database.
//
// WatchChanges sends on ch from the same goroutine as the caller and will deadlock
// if there is no goroutine already listening on it.
//
// Returns ErrNotFound if the request is not reserved at call time.
// If any error is returned, WatchChanges will neither do any operations on ch.
func WatchChanges(reqID string, ch chan<- int) error

// WatchAll watches for all events on the given request.
//
// The WaitGroup will be done after the request is reserved, and the channel
// will send on each state change, then be closed when the request is released.
//
// The channel will send from a new, internal goroutine, which you are not responsible
// for managing.
func WatchAll(reqID string) (*sync.WaitGroup, <-chan int)