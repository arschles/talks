package main

import (
	"net/http"
)

func switchHandler(w http.ResponseWriter, r *http.Request) {
	tmpCh := jsonCh
	jsonCh = base64Ch
	base64Ch = tmpCh
	w.WriteHeader(http.StatusOK)
}
