package main

import (
	"net/http"
)

func sendHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	msgCh := make(chan string)
	streamCh := make(chan chan string)
	streamsCh <- streamCh
	// someone picked up the stream
	for c := range msg {
		msgCh <- c
	}
	// not start sending them messages
}
