package main

type encoderArg struct {
	iface interface{}
	retCh chan []byte
	errCh chan error
}

func newEncoderArg(iface interface{}) encoderArg {
	return encoderArg{
		iface: iface,
		retCh: make(chan []byte),
		errCh: make(chan error),
	}
}
