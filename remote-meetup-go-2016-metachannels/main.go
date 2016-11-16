package main

import (
	"log"
	"net/http"
)

var (
	jsonCh           = make(chan encoderArg)
	base64Ch         = make(chan encoderArg)
	chainedEncoderCh = make(chan chainedEncoderArg)
	stopCh           = make(chan struct{})
)

func main() {
	log.Println("starting 5 encoders")
	for i := 0; i < 5; i++ {
		go jsonEncoder(jsonCh, stopCh)
		go base64Encoder(base64Ch, stopCh)
		go chainedEncoder(chainedEncoderCh)
	}

	/////
	// json and base64 encoders: level I
	/////
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/base64", base64Handler)

	/////
	// switching encoders on the fly: level II
	// we can switch implementations of encodings on the fly, just by switching the channels
	/////
	http.HandleFunc("/switch", switchHandler)

	/////
	// scaling up and down the number of encoders in the worker pool: level III
	/////
	http.HandleFunc("/scale", scaleHandler)

	/////
	// chaining encoders together on the fly: level IV
	/////
	http.HandleFunc("/chained", chainedEncoderHandler)

	// just starting the server
	log.Printf("listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
