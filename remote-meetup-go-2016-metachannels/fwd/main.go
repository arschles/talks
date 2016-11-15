package main

import (
	"net/http"
)

var (
	streamsCollCh = make(chan chan chan string)
)

func main() {
	http.HandleFunc("/send", sendHandler)
	http.HandleFunc("/recv", recvHandler)
	http.ListenAndServe(":8080", nil)
}
