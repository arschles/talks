package main

import (
	"net/http"
)

func switchHandler(w http.ResponseWriter, r *http.Request) {
	// NOTE: this isn't concurrency safe. In production code, use a mutex or manage these channels
	// inside a single goroutine that acts as a multiplexer (kinda like an actor)
	tmpCh := jsonCh
	jsonCh = base64Ch
	base64Ch = tmpCh
	w.WriteHeader(http.StatusOK)
}
