package main

import (
	"net/http"
	"time"
)

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	arg := newEncoderArg(r.URL.Query().Get("val"))

	jsonCh <- arg

	select {
	case <-time.After(100 * time.Millisecond):
		http.Error(w, "failed to encode within 100ms", http.StatusInternalServerError)
	case err := <-arg.errCh:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	case ret := <-arg.retCh:
		w.Write(ret)
	}
}
