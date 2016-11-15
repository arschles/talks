package main

import (
	"net/http"
)

func recvHandler(w http.ResponseWriter, r *http.Request) {
	stream := <-streamsCh
	w.Write([]byte(<-msgCh))
}
