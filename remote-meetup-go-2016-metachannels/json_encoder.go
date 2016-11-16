package main

import (
	"encoding/json"
	"log"
)

func jsonEncoder(argCh <-chan encoderArg, quitCh <-chan struct{}) {
	for {
		select {
		case arg := <-argCh:
			encoded, err := json.Marshal(arg.iface)
			if err != nil {
				arg.errCh <- err
				continue
			}
			arg.retCh <- encoded
		case <-quitCh:
			log.Printf("json encoder stopped")
			return
		}
	}
}
