package main

import (
	"net/http"
	"strconv"
)

// handler to scale up or down the number of encoders in the worker pool by shutting down
// all encoders, then starting up a specified number of them
func scaleHandler(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Query().Get("num"))
	if err != nil {
		http.Error(w, "invalid number", http.StatusInternalServerError)
		return
	}
	// shut down all encoders first
	close(stopCh)
	stopCh = make(chan struct{})
	for i := 0; i < n; i++ {
		go jsonEncoder(jsonCh, stopCh)
		go base64Encoder(base64Ch, stopCh)
	}
	w.WriteHeader(http.StatusOK)
}
