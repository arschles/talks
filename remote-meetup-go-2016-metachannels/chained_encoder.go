package main

import (
	"fmt"
	"time"
)

type chainedEncoderArg struct {
	iface        interface{}
	encoderChans []chan encoderArg
	retCh        chan []byte
	errCh        chan error
}

func newChainedEncoderArg(iface interface{}, encoderChans []chan encoderArg) chainedEncoderArg {
	return chainedEncoderArg{
		iface:        iface,
		encoderChans: encoderChans,
		retCh:        make(chan []byte),
		errCh:        make(chan error),
	}
}

func chainedEncoder(argCh <-chan chainedEncoderArg) {
	for {
		select {
		case arg := <-argCh:
			// convert the argument to bytes, loop through all of the encoders, and string the results
			// through
			res := []byte(fmt.Sprintf("%s", arg.iface))
			totalTimer := time.After(200 * time.Millisecond)
			fail := false
			for i, encoderChan := range arg.encoderChans {
				// create a new arg and send it to this encoder
				eArg := newEncoderArg(res)
				encoderChan <- eArg

				// wait for the encoder to return
				select {
				case ret := <-eArg.retCh:
					// encoder returned successfully, set the current result to the return val
					res = ret
				case err := <-eArg.errCh:
					// encoder returned with error. send it on the original arg's error channel and move on
					arg.errCh <- err
					fail = true
					continue
				case <-totalTimer:
					// encoder didn't return in time. send an appropriate error on the original arg's error
					// channel and move on
					arg.errCh <- fmt.Errorf("we're on encoder %d and 200ms has already elapsed", i)
					fail = true
					continue
				}
			}
			if !fail {
				arg.retCh <- res
			}
		}
	}
}
