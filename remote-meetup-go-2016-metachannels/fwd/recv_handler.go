package main

import (
	"net/http"
)

func recvHandler(w http.ResponseWriter, r *http.Request) {
	discussionCh := <-coordCh
	w.Write([]byte(<-msgCh))
}
