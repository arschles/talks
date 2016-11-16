package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func base64Encoder(argCh <-chan encoderArg, quitCh <-chan struct{}) {
	for {
		select {
		case arg := <-argCh:
			str := fmt.Sprintf("%s", arg.iface)
			encoded := base64.StdEncoding.EncodeToString([]byte(str))
			arg.retCh <- []byte(encoded)
		case <-quitCh:
			log.Printf("base 64 encoder stopped")
			return
		}
	}
}
