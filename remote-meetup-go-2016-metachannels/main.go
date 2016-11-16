package main

import (
	"log"
	"net/http"
)

var (
	jsonCh   = make(chan encoderArg)
	base64Ch = make(chan encoderArg)
	stopCh   = make(chan struct{})
)

func main() {
	log.Println("starting 5 encoders")
	for i := 0; i < 5; i++ {
		go jsonEncoder(jsonCh, stopCh)
		go base64Encoder(base64Ch, stopCh)
	}

	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/base64", base64Handler)

	// we can switch implementations of encodings on the fly, just by switching the channels
	http.HandleFunc("/switch", func(w http.ResponseWriter, r *http.Request) {
		tmpCh := jsonCh
		jsonCh = base64Ch
		base64Ch = tmpCh
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/scale", scaleHandler)

	log.Printf("listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
