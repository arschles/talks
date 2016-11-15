package main

import (
	"net/http"
)

var (
	streamsCh = make(chan chan chan string)
	streamCh  = make(chan chan string)
)

func main() {
	http.HandleFunc("/send", sendHandler)
	http.HandleFunc("/recv", recvHandler)
	http.ListenAndServe(":8080", nil)
}
