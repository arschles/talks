package main

import (
	"net/http"
	"time"
)

func chainedEncoderHandler(w http.ResponseWriter, r *http.Request) {
	encoderStrs := r.URL.Query()["encoder"]
	if len(encoderStrs) < 1 {
		http.Error(w, "no encoders specified", http.StatusBadRequest)
		return
	}
	chans := make([]chan encoderArg, len(encoderStrs))
	for i, encoderStr := range encoderStrs {
		switch encoderStr {
		case "json":
			chans[i] = jsonCh
		case "base64":
			chans[i] = base64Ch
		default:
			http.Error(w, "unsupported encoding type "+encoderStr, http.StatusBadRequest)
			return
		}
	}
	arg := newChainedEncoderArg(r.URL.Query().Get("val"), chans)
	chainedEncoderCh <- arg

	select {
	case <-time.After(200 * time.Millisecond):
		http.Error(w, "failed to encode within 200ms", http.StatusInternalServerError)
	case err := <-arg.errCh:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	case ret := <-arg.retCh:
		w.Write(ret)
	}
}
