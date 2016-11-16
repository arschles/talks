package main

import (
	"log"
	"net/http"
	"time"
)

const numEncoders = 5

func main() {
	jsonCh := make(chan encoderArg)
	base64Ch := make(chan encoderArg)
	stopCh := make(chan struct{})

	log.Printf("starting %d encoders", numEncoders)
	for i := 0; i < numEncoders; i++ {
		go jsonEncoder(jsonCh, stopCh)
		go base64Encoder(base64Ch, stopCh)
	}

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		arg := newEncoderArg(r.URL.Query().Get("val"))
		jsonCh <- arg
		select {
		case err := <-arg.errCh:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		case ret := <-arg.retCh:
			w.Write(ret)
		case <-time.After(100 * time.Millisecond):
			http.Error(w, "failed to encode within 100ms", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/base64", func(w http.ResponseWriter, r *http.Request) {
		arg := newEncoderArg(r.URL.Query().Get("val"))
		base64Ch <- arg
		select {
		case err := <-arg.errCh:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		case ret := <-arg.retCh:
			w.Write(ret)
		case <-time.After(100 * time.Millisecond):
			http.Error(w, "failed to encode within 100ms", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		close(stopCh)
		w.Write([]byte("stopped"))
	})

	log.Printf("listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
